defmodule Pulap.Org.Organization do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "organizations" do
    field :name, :string
    field :description, :string
    field :slug, :string
    field :short_description, :string
    field :created_by, :string
    field :updated_by, :string
    field :owner_id, :binary_id

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(organization, attrs) do
    organization
    |> cast(attrs, [:slug, :name, :short_description, :description, :created_by, :updated_by])
    |> validate_required([:slug, :name, :short_description, :description, :created_by, :updated_by])
    |> unique_constraint(:slug)
  end
end
