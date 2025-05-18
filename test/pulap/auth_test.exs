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

  describe "permissions" do
    alias Pulap.Auth.Permission

    import Pulap.AuthFixtures

    @invalid_attrs %{name: nil, description: nil, slug: nil}

    test "list_permissions/0 returns all permissions" do
      permission = permission_fixture()
      assert Auth.list_permissions() == [permission]
    end

    test "get_permission!/1 returns the permission with given id" do
      permission = permission_fixture()
      assert Auth.get_permission!(permission.id) == permission
    end

    test "create_permission/1 with valid data creates a permission" do
      valid_attrs = %{name: "some name", description: "some description", slug: "some slug"}

      assert {:ok, %Permission{} = permission} = Auth.create_permission(valid_attrs)
      assert permission.name == "some name"
      assert permission.description == "some description"
      assert permission.slug == "some slug"
    end

    test "create_permission/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Auth.create_permission(@invalid_attrs)
    end

    test "update_permission/2 with valid data updates the permission" do
      permission = permission_fixture()
      update_attrs = %{name: "some updated name", description: "some updated description", slug: "some updated slug"}

      assert {:ok, %Permission{} = permission} = Auth.update_permission(permission, update_attrs)
      assert permission.name == "some updated name"
      assert permission.description == "some updated description"
      assert permission.slug == "some updated slug"
    end

    test "update_permission/2 with invalid data returns error changeset" do
      permission = permission_fixture()
      assert {:error, %Ecto.Changeset{}} = Auth.update_permission(permission, @invalid_attrs)
      assert permission == Auth.get_permission!(permission.id)
    end

    test "delete_permission/1 deletes the permission" do
      permission = permission_fixture()
      assert {:ok, %Permission{}} = Auth.delete_permission(permission)
      assert_raise Ecto.NoResultsError, fn -> Auth.get_permission!(permission.id) end
    end

    test "change_permission/1 returns a permission changeset" do
      permission = permission_fixture()
      assert %Ecto.Changeset{} = Auth.change_permission(permission)
    end
  end

  describe "resources" do
    alias Pulap.Auth.Resource

    import Pulap.AuthFixtures

    @invalid_attrs %{name: nil, value: nil, description: nil, kind: nil, slug: nil, created_by: nil, updated_by: nil}

    test "list_resources/0 returns all resources" do
      resource = resource_fixture()
      assert Auth.list_resources() == [resource]
    end

    test "get_resource!/1 returns the resource with given id" do
      resource = resource_fixture()
      assert Auth.get_resource!(resource.id) == resource
    end

    test "create_resource/1 with valid data creates a resource" do
      valid_attrs = %{name: "some name", value: "some value", description: "some description", kind: "some kind", slug: "some slug", created_by: "7488a646-e31f-11e4-aace-600308960662", updated_by: "7488a646-e31f-11e4-aace-600308960662"}

      assert {:ok, %Resource{} = resource} = Auth.create_resource(valid_attrs)
      assert resource.name == "some name"
      assert resource.value == "some value"
      assert resource.description == "some description"
      assert resource.kind == "some kind"
      assert resource.slug == "some slug"
      assert resource.created_by == "7488a646-e31f-11e4-aace-600308960662"
      assert resource.updated_by == "7488a646-e31f-11e4-aace-600308960662"
    end

    test "create_resource/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Auth.create_resource(@invalid_attrs)
    end

    test "update_resource/2 with valid data updates the resource" do
      resource = resource_fixture()
      update_attrs = %{name: "some updated name", value: "some updated value", description: "some updated description", kind: "some updated kind", slug: "some updated slug", created_by: "7488a646-e31f-11e4-aace-600308960668", updated_by: "7488a646-e31f-11e4-aace-600308960668"}

      assert {:ok, %Resource{} = resource} = Auth.update_resource(resource, update_attrs)
      assert resource.name == "some updated name"
      assert resource.value == "some updated value"
      assert resource.description == "some updated description"
      assert resource.kind == "some updated kind"
      assert resource.slug == "some updated slug"
      assert resource.created_by == "7488a646-e31f-11e4-aace-600308960668"
      assert resource.updated_by == "7488a646-e31f-11e4-aace-600308960668"
    end

    test "update_resource/2 with invalid data returns error changeset" do
      resource = resource_fixture()
      assert {:error, %Ecto.Changeset{}} = Auth.update_resource(resource, @invalid_attrs)
      assert resource == Auth.get_resource!(resource.id)
    end

    test "delete_resource/1 deletes the resource" do
      resource = resource_fixture()
      assert {:ok, %Resource{}} = Auth.delete_resource(resource)
      assert_raise Ecto.NoResultsError, fn -> Auth.get_resource!(resource.id) end
    end

    test "change_resource/1 returns a resource changeset" do
      resource = resource_fixture()
      assert %Ecto.Changeset{} = Auth.change_resource(resource)
    end
  end

  describe "teams" do
    alias Pulap.Auth.Team

    import Pulap.AuthFixtures

    @invalid_attrs %{name: nil, description: nil, kind: nil, slug: nil, inserted_by: nil, updated_by: nil}

    test "list_teams/0 returns all teams" do
      team = team_fixture()
      assert Auth.list_teams() == [team]
    end

    test "get_team!/1 returns the team with given id" do
      team = team_fixture()
      assert Auth.get_team!(team.id) == team
    end

    test "create_team/1 with valid data creates a team" do
      valid_attrs = %{name: "some name", description: "some description", kind: "some kind", slug: "some slug", inserted_by: 42, updated_by: 42}

      assert {:ok, %Team{} = team} = Auth.create_team(valid_attrs)
      assert team.name == "some name"
      assert team.description == "some description"
      assert team.kind == "some kind"
      assert team.slug == "some slug"
      assert team.inserted_by == 42
      assert team.updated_by == 42
    end

    test "create_team/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Auth.create_team(@invalid_attrs)
    end

    test "update_team/2 with valid data updates the team" do
      team = team_fixture()
      update_attrs = %{name: "some updated name", description: "some updated description", kind: "some updated kind", slug: "some updated slug", inserted_by: 43, updated_by: 43}

      assert {:ok, %Team{} = team} = Auth.update_team(team, update_attrs)
      assert team.name == "some updated name"
      assert team.description == "some updated description"
      assert team.kind == "some updated kind"
      assert team.slug == "some updated slug"
      assert team.inserted_by == 43
      assert team.updated_by == 43
    end

    test "update_team/2 with invalid data returns error changeset" do
      team = team_fixture()
      assert {:error, %Ecto.Changeset{}} = Auth.update_team(team, @invalid_attrs)
      assert team == Auth.get_team!(team.id)
    end

    test "delete_team/1 deletes the team" do
      team = team_fixture()
      assert {:ok, %Team{}} = Auth.delete_team(team)
      assert_raise Ecto.NoResultsError, fn -> Auth.get_team!(team.id) end
    end

    test "change_team/1 returns a team changeset" do
      team = team_fixture()
      assert %Ecto.Changeset{} = Auth.change_team(team)
    end
  end
end
