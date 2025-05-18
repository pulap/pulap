defmodule PulapWeb.TeamController do
  use PulapWeb, :controller

  plug :put_layout, html: {PulapWeb.Layouts, :auth}

  alias Pulap.Auth
  alias Pulap.Org.Team
  alias PulapWeb.AuditHelpers

  def index(conn, _params) do
    IO.inspect(:index_start, label: "[DEBUG] TeamController.index")
    teams = Auth.list_teams()
    IO.inspect(teams, label: "[DEBUG] Loaded teams")
    render(conn, :index, teams: teams)
  end

  def new(conn, _params) do
    IO.inspect(:new_start, label: "[DEBUG] TeamController.new")
    changeset = Auth.change_team(%Team{})
    IO.inspect(changeset, label: "[DEBUG] New changeset")
    render(conn, :new, changeset: changeset)
  end

  def create(conn, %{"team" => team_params}) do
    IO.inspect(team_params, label: "[DEBUG] Incoming team_params")
    params = AuditHelpers.maybe_put_created_by(team_params, conn)
    params = Map.put(params, "updated_by", params["created_by"])
    # Always use the first organization in the database
    organization = Pulap.Repo.one(Pulap.Org.Organization)
    params = Map.put(params, "organization_id", organization && organization.id)
    IO.inspect(params, label: "[DEBUG] Params after maybe_put_created_by, updated_by, and organization_id")
    case Auth.create_team(params) do
      {:ok, team} ->
        IO.inspect(team, label: "[DEBUG] Created team")
        conn
        |> put_flash(:info, "Team created successfully.")
        |> redirect(to: ~p"/teams/#{team}")

      {:error, %Ecto.Changeset{} = changeset} ->
        IO.inspect(changeset, label: "[DEBUG] Team creation failed changeset")
        render(conn, :new, changeset: changeset)
    end
  end

  def show(conn, %{"id" => id}) do
    IO.inspect(id, label: "[DEBUG] Show team id")
    team = Auth.get_team!(id)
    IO.inspect(team, label: "[DEBUG] Loaded team for show")
    render(conn, :show, team: team)
  end

  def edit(conn, %{"id" => id}) do
    IO.inspect(id, label: "[DEBUG] Edit team id")
    team = Auth.get_team!(id)
    IO.inspect(team, label: "[DEBUG] Loaded team for edit")
    changeset = Auth.change_team(team)
    IO.inspect(changeset, label: "[DEBUG] Edit changeset")
    render(conn, :edit, team: team, changeset: changeset)
  end

  def update(conn, %{"id" => id, "team" => team_params}) do
    IO.inspect({id, team_params}, label: "[DEBUG] Update team id and params")
    team = Auth.get_team!(id)
    IO.inspect(team, label: "[DEBUG] Loaded team for update")
    params = AuditHelpers.maybe_put_updated_by(team_params, conn)
    IO.inspect(params, label: "[DEBUG] Params after maybe_put_updated_by")
    case Auth.update_team(team, params) do
      {:ok, team} ->
        IO.inspect(team, label: "[DEBUG] Updated team")
        conn
        |> put_flash(:info, "Team updated successfully.")
        |> redirect(to: ~p"/teams/#{team}")

      {:error, %Ecto.Changeset{} = changeset} ->
        IO.inspect(changeset, label: "[DEBUG] Team update failed changeset")
        render(conn, :edit, team: team, changeset: changeset)
    end
  end

  def delete(conn, %{"id" => id}) do
    IO.inspect(id, label: "[DEBUG] Delete team id")
    team = Auth.get_team!(id)
    IO.inspect(team, label: "[DEBUG] Loaded team for delete")
    {:ok, _team} = Auth.delete_team(team)
    IO.puts("[DEBUG] Team deleted")
    conn
    |> put_flash(:info, "Team deleted successfully.")
    |> redirect(to: ~p"/teams")
  end
end
