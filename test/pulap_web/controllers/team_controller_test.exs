defmodule PulapWeb.TeamControllerTest do
  use PulapWeb.ConnCase

  import Pulap.AuthFixtures

  @create_attrs %{name: "some name", description: "some description", kind: "some kind", slug: "some slug", inserted_by: 42, updated_by: 42}
  @update_attrs %{name: "some updated name", description: "some updated description", kind: "some updated kind", slug: "some updated slug", inserted_by: 43, updated_by: 43}
  @invalid_attrs %{name: nil, description: nil, kind: nil, slug: nil, inserted_by: nil, updated_by: nil}

  describe "index" do
    test "lists all teams", %{conn: conn} do
      conn = get(conn, ~p"/teams")
      assert html_response(conn, 200) =~ "Listing Teams"
    end
  end

  describe "new team" do
    test "renders form", %{conn: conn} do
      conn = get(conn, ~p"/teams/new")
      assert html_response(conn, 200) =~ "New Team"
    end
  end

  describe "create team" do
    test "redirects to show when data is valid", %{conn: conn} do
      conn = post(conn, ~p"/teams", team: @create_attrs)

      assert %{id: id} = redirected_params(conn)
      assert redirected_to(conn) == ~p"/teams/#{id}"

      conn = get(conn, ~p"/teams/#{id}")
      assert html_response(conn, 200) =~ "Team #{id}"
    end

    test "renders errors when data is invalid", %{conn: conn} do
      conn = post(conn, ~p"/teams", team: @invalid_attrs)
      assert html_response(conn, 200) =~ "New Team"
    end
  end

  describe "edit team" do
    setup [:create_team]

    test "renders form for editing chosen team", %{conn: conn, team: team} do
      conn = get(conn, ~p"/teams/#{team}/edit")
      assert html_response(conn, 200) =~ "Edit Team"
    end
  end

  describe "update team" do
    setup [:create_team]

    test "redirects when data is valid", %{conn: conn, team: team} do
      conn = put(conn, ~p"/teams/#{team}", team: @update_attrs)
      assert redirected_to(conn) == ~p"/teams/#{team}"

      conn = get(conn, ~p"/teams/#{team}")
      assert html_response(conn, 200) =~ "some updated name"
    end

    test "renders errors when data is invalid", %{conn: conn, team: team} do
      conn = put(conn, ~p"/teams/#{team}", team: @invalid_attrs)
      assert html_response(conn, 200) =~ "Edit Team"
    end
  end

  describe "delete team" do
    setup [:create_team]

    test "deletes chosen team", %{conn: conn, team: team} do
      conn = delete(conn, ~p"/teams/#{team}")
      assert redirected_to(conn) == ~p"/teams"

      assert_error_sent 404, fn ->
        get(conn, ~p"/teams/#{team}")
      end
    end
  end

  defp create_team(_) do
    team = team_fixture()
    %{team: team}
  end
end
