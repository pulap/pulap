defmodule Pulap.Dict.Entry do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "entries" do
    field :active, :boolean, default: false
    field :label, :string
    field :position, :integer
    field :value, :string
    field :created_by, Ecto.UUID
    field :updated_by, Ecto.UUID
    field :dictionary_id, :binary_id

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(entry, attrs) do
    entry
    |> cast(attrs, [:value, :label, :position, :active])
    |> validate_required([:value, :label, :position, :active])
  end
end
