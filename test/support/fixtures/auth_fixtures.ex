defmodule Pulap.AuthFixtures do
  @moduledoc """
  This module defines test helpers for creating
  entities via the `Pulap.Auth` context.
  """

  @doc """
  Generate a role.
  """
  def role_fixture(attrs \\ %{}) do
    {:ok, role} =
      attrs
      |> Enum.into(%{
        name: "some name",
        description: "some description"
      })
      |> Pulap.Auth.create_role()

    role
  end

  @doc """
  Generate a permission.
  """
  def permission_fixture(attrs \\ %{}) do
    {:ok, permission} =
      attrs
      |> Enum.into(%{
        description: "some description",
        name: "some name",
        slug: "some slug"
      })
      |> Pulap.Auth.create_permission()

    permission
  end

  @doc """
  Generate a resource.
  """
  def resource_fixture(attrs \\ %{}) do
    {:ok, resource} =
      attrs
      |> Enum.into(%{
        name: "some name",
        value: "some value",
        description: "some description",
        kind: "some kind",
        slug: "some-slug-#{System.unique_integer()}",
        created_by: "7488a646-e31f-11e4-aace-600308960662",
        updated_by: "7488a646-e31f-11e4-aace-600308960662"
      })
      |> Pulap.Auth.create_resource()

    resource
  end

  @doc """
  Generate a team.
  """
  def team_fixture(attrs \\ %{}) do
    attrs =
      attrs
      |> Enum.map(fn
        {k, v} when is_binary(k) -> {String.to_atom(k), v}
        pair -> pair
      end)
      |> Enum.into(%{})

    org =
      Map.get(attrs, :organization) || Map.get(attrs, "organization") || organization_fixture()

    attrs = Map.put(attrs, :organization_id, org.id)

    {:ok, team} =
      attrs
      |> Enum.into(%{
        description: "some description",
        created_by: 42,
        kind: "some kind",
        name: "some name",
        slug: "some slug",
        updated_by: 42
      })
      |> Pulap.Auth.create_team()

    team
  end

  @doc """
  Generate an organization.
  """
  def organization_fixture(attrs \\ %{}) do
    {:ok, organization} =
      attrs
      |> Enum.into(%{
        name: "Test Org",
        slug: "test-org-#{System.unique_integer()}",
        description: "A test organization"
      })
      |> Pulap.Auth.create_organization()

    organization
  end
end
