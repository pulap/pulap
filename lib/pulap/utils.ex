defmodule Pulap.Utils do
  @moduledoc """
  Utility functions for Pulap schemas.
  """

  import Ecto.Changeset

  @doc """
  Generates a unique short code by taking the last segment of a UUID
  """
  def generate_short_code do
    Ecto.UUID.generate() |> String.split("-") |> List.last()
  end

  @doc """
  Generates a slug by combining the slugified source text with a short code.
  """
  def generate_slug(text, short_code) when is_binary(text) and is_binary(short_code) do
    slugify(text) <> "-" <> short_code
  end

  @doc """
  Puts a short_code into the changeset.
  """
  def put_short_code(%Ecto.Changeset{} = changeset, field) when is_atom(field) do
    if get_change(changeset, field) do
      changeset
    else
      put_change(changeset, field, generate_short_code())
    end
  end

  @doc """
  Converts a string into a URL-friendly slug.
  """
  def slugify(str) when is_binary(str) do
    str
    |> String.downcase()
    |> String.replace(~r/[^\w-]+/u, "-")
    |> String.replace(~r/-+/, "-")
    |> String.trim("-")
  end

  @doc """
  Gets a dynamic slug for a struct using its SlugSource and short_code.
  """
  def get_slug(struct) do
    source = Pulap.SlugSource.source_for_slug(struct)
    generate_slug(source, struct.short_code)
  end
end
