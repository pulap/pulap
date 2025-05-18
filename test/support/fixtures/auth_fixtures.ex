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
        slug: "some slug",
        status: "some status"
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
        created_by: "7488a646-e31f-11e4-aace-600308960662",
        description: "some description",
        kind: "some kind",
        name: "some name",
        slug: "some slug",
        updated_by: "7488a646-e31f-11e4-aace-600308960662",
        value: "some value"
      })
      |> Pulap.Auth.create_resource()

    resource
  end

  @doc """
  Generate a team.
  """
  def team_fixture(attrs \\ %{}) do
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
      |> Pulap.Org.create_team()

    team
  end
end
