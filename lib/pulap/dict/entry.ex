defmodule Pulap.Dict.Entry do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "dictionary_entries" do
    field :short_code, :string
    field :active, :boolean, default: true
    field :label, :string
    field :description, :string
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
    |> put_short_code(:short_code)
  end

  def slug(%__MODULE__{} = entry) do
    Pulap.Utils.get_slug(entry)
  end
end

defimpl Pulap.SlugSource, for: Pulap.Dict.Entry do
  def source_for_slug(%Pulap.Dict.Entry{value: value}) do
    value
  end
end
