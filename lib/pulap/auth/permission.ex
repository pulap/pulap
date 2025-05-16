defmodule Pulap.Auth.Permission do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "permissions" do
    field :name, :string
    field :description, :string
    field :slug, :string
    field :created_by, :binary_id
    field :updated_by, :binary_id

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(permission, attrs) do
    permission
    |> cast(attrs, [:slug, :name, :description])
    |> validate_required([:slug, :name, :description])
  end
end
