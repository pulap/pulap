defmodule Pulap.Utils do
  @moduledoc """
  Utility functions for Pulap schemas.
  """

  import Ecto.Changeset

  @doc """
  Puts a slug generated from the :name field and a UUID segment into the changeset.
  """
  def put_slug(changeset, field \\ :slug) do
    if name = get_field(changeset, :name) do
      uuid = Ecto.UUID.generate()
      last_segment = uuid |> String.split("-") |> List.last()

      slug =
        name
        |> String.downcase()
        |> String.replace(~r/[^a-z0-9]+/u, "-")
        |> String.trim("-")
        |> Kernel.<>("-" <> last_segment)

      put_change(changeset, field, slug)
    else
      changeset
    end
  end
end
