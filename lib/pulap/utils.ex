defmodule Pulap.Utils do
  @moduledoc """
  Utility functions for Pulap schemas.
  """

  import Ecto.Changeset

  @doc """
  Puts a slug generated as a UUID segment into the changeset.
  """
  def put_slug(changeset, field \\ :slug, _source_field \\ :name) do
    uuid = Ecto.UUID.generate()
    slug = uuid |> String.split("-") |> List.last()
    put_change(changeset, field, slug)
  end
end
