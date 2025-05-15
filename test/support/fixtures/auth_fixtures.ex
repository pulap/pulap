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
end
