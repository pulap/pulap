defmodule Pulap.OrgFixtures do
  @moduledoc """
  This module defines test helpers for creating
  entities via the `Pulap.Org` context.
  """

  @doc """
  Generate a unique organization slug.
  """
  def unique_organization_slug, do: "some slug#{System.unique_integer([:positive])}"

  @doc """
  Generate a organization.
  """
  def organization_fixture(attrs \\ %{}) do
    {:ok, organization} =
      attrs
      |> Enum.into(%{
        created_by: "some created_by",
        description: "some description",
        name: "some name",
        short_description: "some short_description",
        slug: unique_organization_slug(),
        updated_by: "some updated_by"
      })
      |> Pulap.Org.create_organization()

    organization
  end
end
