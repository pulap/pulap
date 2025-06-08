defmodule Pulap.Dict.Dictionary do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "dictionaries" do
    field :short_code, :string
    field :label, :string
    field :description, :string
    field :created_by, Ecto.UUID
    field :updated_by, Ecto.UUID

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(dictionary, attrs) do
    dictionary
    |> cast(attrs, [:short_code, :label, :description, :created_by, :updated_by])
    |> validate_required([:short_code, :label, :description])
    |> unique_constraint(:short_code)
  end
end
