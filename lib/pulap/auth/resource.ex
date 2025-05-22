defmodule Pulap.Auth.Resource do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "resources" do
    field :name, :string
    field :value, :string
    field :description, :string
    field :kind, :string
    field :slug, :string
    field :created_by, Ecto.UUID
    field :updated_by, Ecto.UUID

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(resource, attrs) do
    resource
    |> cast(attrs, [:name, :value, :description, :kind])
    |> put_slug()
    |> validate_required([:name, :value, :description, :kind, :slug])
    |> unique_constraint(:slug)
  end
end
