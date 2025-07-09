defmodule Pulap.Set.Option do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "options" do
    field :short_code, :string
    field :key, :string
    field :label, :string
    field :description, :string
    field :value, :string
    field :order, :integer, default: 0
    field :active, :boolean, default: true
    field :created_by, Ecto.UUID
    field :updated_by, Ecto.UUID
    field :parent_id, :binary_id

    belongs_to :set, Pulap.Set.Set, type: :binary_id

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(option, attrs) do
    option
    |> cast(attrs, [
      :id,
      :key,
      :label,
      :description,
      :value,
      :order,
      :active,
      :set_id,
      :parent_id,
      :created_by,
      :updated_by
    ])
    |> put_short_code(:short_code)
    |> validate_required([:key, :label, :value, :set_id])
    |> unique_constraint(:short_code, name: :options_short_code_set_id_index)
    |> unique_constraint(:key, name: :options_key_set_id_index)
  end

  end

defimpl Pulap.SlugSource, for: Pulap.Set.Option do
  def source_for_slug(%Pulap.Set.Option{value: value}) do
    value
  end
end
