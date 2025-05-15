defmodule Pulap.AuthTest do
  use Pulap.DataCase

  alias Pulap.Auth

  describe "roles" do
    alias Pulap.Auth.Role

    import Pulap.AuthFixtures

    @invalid_attrs %{name: nil, status: nil, slug: nil}

    test "list_roles/0 returns all roles" do
      role = role_fixture()
      assert Auth.list_roles() == [role]
    end

    test "get_role!/1 returns the role with given id" do
      role = role_fixture()
      assert Auth.get_role!(role.id) == role
    end

    test "create_role/1 with valid data creates a role" do
      valid_attrs = %{name: "some name", status: "some status", slug: "some slug"}

      assert {:ok, %Role{} = role} = Auth.create_role(valid_attrs)
      assert role.name == "some name"
      assert role.status == "some status"
      assert role.slug == "some slug"
    end

    test "create_role/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Auth.create_role(@invalid_attrs)
    end

    test "update_role/2 with valid data updates the role" do
      role = role_fixture()
      update_attrs = %{name: "some updated name", status: "some updated status", slug: "some updated slug"}

      assert {:ok, %Role{} = role} = Auth.update_role(role, update_attrs)
      assert role.name == "some updated name"
      assert role.status == "some updated status"
      assert role.slug == "some updated slug"
    end

    test "update_role/2 with invalid data returns error changeset" do
      role = role_fixture()
      assert {:error, %Ecto.Changeset{}} = Auth.update_role(role, @invalid_attrs)
      assert role == Auth.get_role!(role.id)
    end

    test "delete_role/1 deletes the role" do
      role = role_fixture()
      assert {:ok, %Role{}} = Auth.delete_role(role)
      assert_raise Ecto.NoResultsError, fn -> Auth.get_role!(role.id) end
    end

    test "change_role/1 returns a role changeset" do
      role = role_fixture()
      assert %Ecto.Changeset{} = Auth.change_role(role)
    end
  end
end
