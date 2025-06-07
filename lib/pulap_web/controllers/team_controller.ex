defmodule PulapWeb.TeamController do
  use PulapWeb, :controller

  plug :put_layout, html: {PulapWeb.Layouts, :auth}

  alias Pulap.{Auth, Org}
  alias Pulap.Org.Team
  alias PulapWeb.AuditHelpers
  import Ecto.Query

  def index(conn, _params) do
    teams = Auth.list_teams()
    render(conn, :index, teams: teams)
  end

  def new(conn, _params) do
    changeset = Auth.change_team(%Team{})
    render(conn, :new, changeset: changeset)
  end

  def create(conn, %{"team" => team_params}) do
    params = AuditHelpers.maybe_put_created_by(team_params, conn)
    params = Map.put(params, "updated_by", params["created_by"])
    # Always use the first organization in the database
    organization = List.first(Org.list_organizations())
    params = Map.put(params, "organization_id", organization && organization.id)

    case Auth.create_team(params) do
      {:ok, team} ->
        conn
        |> put_flash(:info, "Team created successfully.")
        |> redirect(to: ~p"/teams/#{team}")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :new, changeset: changeset)
    end
  end

  def show(conn, %{"id" => id}) do
    team = Auth.get_team!(id)
    render(conn, :show, team: team)
  end

  def edit(conn, %{"id" => id}) do
    team = Auth.get_team!(id)
    changeset = Auth.change_team(team)
    render(conn, :edit, team: team, changeset: changeset)
  end

  def update(conn, %{"id" => id, "team" => team_params}) do
    team = Auth.get_team!(id)
    params = AuditHelpers.maybe_put_updated_by(team_params, conn)

    case Auth.update_team(team, params) do
      {:ok, team} ->
        conn
        |> put_flash(:info, "Team updated successfully.")
        |> redirect(to: ~p"/teams/#{team}")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :edit, team: team, changeset: changeset)
    end
  end

  def delete(conn, %{"id" => id}) do
    team = Auth.get_team!(id)
    {:ok, _team} = Auth.delete_team(team)

    conn
    |> put_flash(:info, "Team deleted successfully.")
    |> redirect(to: ~p"/teams")
  end

  def members(conn, %{"team_id" => team_id}) do
    team = Auth.get_team!(team_id)
    # Single query to get all users and their assignment status
    users_with_status =
      from(u in Pulap.Accounts.User,
        left_join: tm in Pulap.Org.TeamMembership,
        on: tm.user_id == u.id and tm.team_id == ^team_id,
        select: %{user: u, assigned: not is_nil(tm.id)}
      )
      |> Pulap.Repo.all()

    assigned = Enum.filter(users_with_status, & &1.assigned) |> Enum.map(& &1.user)
    unassigned = Enum.reject(users_with_status, & &1.assigned) |> Enum.map(& &1.user)
    render(conn, :members, team: team, assigned: assigned, unassigned: unassigned)
  end

  def assign_member(conn, %{"team_id" => team_id, "user_id" => user_id}) do
    attrs = %{team_id: team_id, user_id: user_id, relation_type: "direct"}

    case Pulap.Org.TeamMembership.changeset(%Pulap.Org.TeamMembership{}, attrs)
         |> Pulap.Repo.insert() do
      {:ok, _} ->
        conn
        |> put_flash(:info, "Member assigned successfully.")
        |> redirect(to: ~p"/teams/#{team_id}/members")

      {:error, _} ->
        conn
        |> put_flash(:error, "Could not assign member.")
        |> redirect(to: ~p"/teams/#{team_id}/members")
    end
  end

  def delete_member(conn, %{"team_id" => team_id, "id" => user_id}) do
    from(tm in Pulap.Org.TeamMembership, where: tm.team_id == ^team_id and tm.user_id == ^user_id)
    |> Pulap.Repo.delete_all()

    conn
    |> put_flash(:info, "Member revoked successfully.")
    |> redirect(to: ~p"/teams/#{team_id}/members")
  end

  def member_roles(conn, %{"team_id" => team_id, "id" => user_id}) do
    team = Auth.get_team!(team_id)
    user = Pulap.Accounts.get_user!(user_id)

    # Get contextual roles assigned to user in this team context
    assigned_roles_query =
      from(r in Pulap.Auth.Role,
        join: ur in "users_roles",
        on: ur.role_id == r.id,
        where: r.contextual == true,
        where: ur.user_id == ^user_id,
        where: ur.context_type == "team",
        where: ur.context_id == ^team_id
      )

    # Get unassigned contextual roles that aren't assigned in this team context
    unassigned_roles_query =
      from(r in Pulap.Auth.Role,
        where: r.contextual == true,
        left_join: ur in "users_roles",
        on:
          ur.role_id == r.id and ur.user_id == ^user_id and ur.context_type == "team" and
            ur.context_id == ^team_id,
        where: is_nil(ur.role_id)
      )

    assigned_roles = Pulap.Repo.all(assigned_roles_query)
    unassigned_roles = Pulap.Repo.all(unassigned_roles_query)

    render(conn, :member_roles,
      team: team,
      user: user,
      assigned_roles: assigned_roles,
      unassigned_roles: unassigned_roles
    )
  end

  def assign_member_role(conn, %{"team_id" => team_id, "id" => user_id, "role_id" => role_id}) do
    attrs = %{
      user_id: user_id,
      role_id: role_id,
      context_type: "team",
      context_id: team_id
    }

    case Pulap.Repo.insert_all("users_roles", [attrs], on_conflict: :nothing) do
      {1, _} ->
        conn
        |> put_flash(:info, "Role assigned successfully.")
        |> redirect(to: ~p"/teams/#{team_id}/members/#{user_id}/roles")

      {0, _} ->
        conn
        |> put_flash(:error, "Role was already assigned.")
        |> redirect(to: ~p"/teams/#{team_id}/members/#{user_id}/roles")
    end
  end

  def revoke_member_role(conn, %{"team_id" => team_id, "id" => user_id, "role_id" => role_id}) do
    {count, _} =
      from(ur in "users_roles",
        where: ur.user_id == ^user_id,
        where: ur.role_id == ^role_id,
        where: ur.context_type == "team",
        where: ur.context_id == ^team_id
      )
      |> Pulap.Repo.delete_all()

    if count > 0 do
      conn
      |> put_flash(:info, "Role revoked successfully.")
      |> redirect(to: ~p"/teams/#{team_id}/members/#{user_id}/roles")
    else
      conn
      |> put_flash(:error, "Role was not assigned.")
      |> redirect(to: ~p"/teams/#{team_id}/members/#{user_id}/roles")
    end
  end
end
