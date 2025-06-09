defmodule Pulap.Set.Set do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "sets" do
    field :short_code, :string
    field :key, :string
    field :label, :string
    field :description, :string
    field :created_by, :binary_id
    field :updated_by, :binary_id

    has_many :options, Pulap.Set.Option

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(set, attrs) do
    set
    |> cast(attrs, [:key, :label, :description])
    |> put_short_code(:short_code)
    |> validate_required([:key, :label, :description, :short_code])
    |> unique_constraint(:key)
  end
end
