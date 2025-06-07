defmodule Pulap.Dict.Entry do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "dictionary_entries" do
    field :active, :boolean, default: true
    field :label, :string
    field :description, :string
    field :short_code, :string
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
    |> cast(attrs, [:value, :label, :description, :order, :active, :dictionary_id])
    |> validate_required([:value, :label, :dictionary_id])
    |> put_slug(:short_code, :value)
    |> unique_constraint([:short_code, :dictionary_id])
  end
end
