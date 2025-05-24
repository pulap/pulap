defmodule Pulap.Dict.Dictionary do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "dictionaries" do
    field :label, :string
    field :description, :string
    field :slug, :string
    field :created_by, Ecto.UUID
    field :updated_by, Ecto.UUID

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(dictionary, attrs) do
    dictionary
    |> cast(attrs, [:slug, :label, :description])
    |> validate_required([:slug, :label, :description])
    |> unique_constraint(:slug)
  end
end
