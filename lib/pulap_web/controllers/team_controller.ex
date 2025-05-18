defmodule PulapWeb.TeamController do
  use PulapWeb, :controller

  plug :put_layout, {PulapWeb.Layouts, :auth}

  alias Pulap.Auth
  alias Pulap.Auth.Team
  alias PulapWeb.AuditHelpers

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
end
