defmodule PulapWeb.ResourceControllerTest do
  use PulapWeb.ConnCase

  import Pulap.AuthFixtures

  @invalid_attrs %{
    name: nil,
    value: nil,
    description: nil,
    kind: nil,
    slug: nil,
    created_by: nil,
    updated_by: nil
  }

  setup :register_and_log_in_user

  describe "index" do
    # test "lists all resources", %{conn: conn} do
    #   conn = get(conn, ~p"/resources")
    #   assert html_response(conn, 200) =~ "Listing Resources"
    # end
  end

  describe "new resource" do
    # test "renders form", %{conn: conn} do
    #   conn = get(conn, ~p"/resources/new")
    #   assert html_response(conn, 200) =~ "New Resource"
    # end
  end

  describe "create resource" do
    # test "redirects to show when data is valid", %{conn: conn} do
    #   conn = post(conn, ~p"/resources", resource: @create_attrs)

    #   assert %{id: id} = redirected_params(conn)
    #   assert redirected_to(conn) == ~p"/resources/#{id}"

    #   conn = get(conn, ~p"/resources/#{id}")
    #   assert html_response(conn, 200) =~ "Resource #{id}"
    # end

    test "renders errors when data is invalid", %{conn: conn} do
      conn = post(conn, ~p"/resources", resource: @invalid_attrs)
      assert html_response(conn, 200) =~ "New Resource"
    end
  end

  describe "edit resource" do
    setup [:create_resource]

    # test "renders form for editing chosen resource", %{conn: conn, resource: resource} do
    #   conn = get(conn, ~p"/resources/#{resource}/edit")
    #   assert html_response(conn, 200) =~ "Edit Resource"
    # end
  end

  describe "update resource" do
    setup [:create_resource]

    # test "redirects when data is valid", %{conn: conn, resource: resource} do
    #   conn = put(conn, ~p"/resources/#{resource}", resource: @update_attrs)
    #   assert redirected_to(conn) == ~p"/resources/#{resource}"

    #   conn = get(conn, ~p"/resources/#{resource}")
    #   assert html_response(conn, 200) =~ "some updated name"
    # end

    test "renders errors when data is invalid", %{conn: conn, resource: resource} do
      conn = put(conn, ~p"/resources/#{resource}", resource: @invalid_attrs)
      assert html_response(conn, 200) =~ "Edit Resource"
    end
  end

  describe "delete resource" do
    setup [:create_resource]

    test "deletes chosen resource", %{conn: conn, resource: resource} do
      conn = delete(conn, ~p"/resources/#{resource}")
      assert redirected_to(conn) == ~p"/resources"

      assert_error_sent 404, fn ->
        get(conn, ~p"/resources/#{resource}")
      end
    end
  end

  defp create_resource(_) do
    resource = resource_fixture()
    %{resource: resource}
  end
end
