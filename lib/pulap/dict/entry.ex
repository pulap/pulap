defmodule Pulap.Dict.Entry do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "dictionary_entries" do
    field :active, :boolean, default: true
    field :label, :string
    field :description, :string
    field :slug, :string
    field :value, :string
    field :order, :integer, default: 0
    field :created_by, Ecto.UUID
    field :updated_by, Ecto.UUID
    belongs_to :dictionary, Pulap.Dict.Dictionary, type: :binary_id

    timestamps()
  end

  @doc false
  def changeset(entry, attrs) do
    entry
    |> cast(attrs, [:value, :label, :description, :slug, :order, :active, :dictionary_id])
    |> validate_required([:value, :label, :slug, :dictionary_id])
    |> unique_constraint([:slug, :dictionary_id])
  end
end
