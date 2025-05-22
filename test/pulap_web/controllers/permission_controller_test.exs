defmodule PulapWeb.PermissionControllerTest do
  use PulapWeb.ConnCase

  import Pulap.AuthFixtures

  setup :register_and_log_in_user

  describe "index" do
    # test "lists all permissions", %{conn: conn} do
    #   conn = get(conn, ~p"/permissions")
    #   assert html_response(conn, 200) =~ "Listing Permissions"
    # end
  end

  describe "new permission" do
    # test "renders form", %{conn: conn} do
    #   conn = get(conn, ~p"/permissions/new")
    #   assert html_response(conn, 200) =~ "New Permission"
    # end
  end

  describe "create permission" do
    # test "redirects to show when data is valid", %{conn: conn} do
    #   conn = post(conn, ~p"/permissions", permission: @create_attrs)

    #   assert %{id: id} = redirected_params(conn)
    #   assert redirected_to(conn) == ~p"/permissions/#{id}"

    #   conn = get(conn, ~p"/permissions/#{id}")
    #   assert html_response(conn, 200) =~ "Permission #{id}"
    # end

    # test "renders errors when data is invalid", %{conn: conn} do
    #   conn = post(conn, ~p"/permissions", permission: @invalid_attrs)
    #   assert html_response(conn, 200) =~ "New Permission"
    # end
  end

  describe "edit permission" do
    setup [:create_permission]

    # test "renders form for editing chosen permission", %{conn: conn, permission: permission} do
    #   conn = get(conn, ~p"/permissions/#{permission}/edit")
    #   assert html_response(conn, 200) =~ "Edit Permission"
    # end
  end

  describe "update permission" do
    setup [:create_permission]

    # test "redirects when data is valid", %{conn: conn, permission: permission} do
    #   conn = put(conn, ~p"/permissions/#{permission}", permission: @update_attrs)
    #   assert redirected_to(conn) == ~p"/permissions/#{permission}"

    #   conn = get(conn, ~p"/permissions/#{permission}")
    #   assert html_response(conn, 200) =~ "some updated name"
    # end

    # test "renders errors when data is invalid", %{conn: conn, permission: permission} do
    #   conn = put(conn, ~p"/permissions/#{permission}", permission: @invalid_attrs)
    #   assert html_response(conn, 200) =~ "Edit Permission"
    # end
  end

  describe "delete permission" do
    setup [:create_permission]

    # test "deletes chosen permission", %{conn: conn, permission: permission} do
    #   conn = delete(conn, ~p"/permissions/#{permission}")
    #   assert redirected_to(conn) == ~p"/permissions"

    #   assert_error_sent 404, fn ->
    #     get(conn, ~p"/permissions/#{permission}")
    #   end
    # end
  end

  defp create_permission(_) do
    permission = permission_fixture()
    %{permission: permission}
  end
end
