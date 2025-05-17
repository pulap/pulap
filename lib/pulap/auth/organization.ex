defmodule Pulap.Auth.Organization do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "organizations" do
    field :slug, :string
    field :name, :string
    field :description, :string
    field :created_by, :binary_id
    field :updated_by, :binary_id
    belongs_to :owner, Pulap.Accounts.User, type: :binary_id
    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(organization, attrs) do
    organization
    |> cast(attrs, [:slug, :name, :description, :created_by, :updated_by, :owner_id])
    |> validate_required([:name])
    |> unique_constraint(:slug)
  end
end