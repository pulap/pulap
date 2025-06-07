defmodule Pulap.Org.Organization do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "organizations" do
    field :short_code, :string
    field :name, :string
    field :short_description, :string
    field :description, :string
    field :created_by, :binary_id
    field :updated_by, :binary_id

    many_to_many :owners, Pulap.Accounts.User,
      join_through: Pulap.Org.OrganizationOwner,
      join_keys: [organization_id: :id, user_id: :id]

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(organization, attrs) do
    organization
    |> cast(attrs, [:short_code, :name, :short_description, :description])
    |> put_slug(:short_code)
    |> validate_required([:name])
    |> unique_constraint(:short_code)
  end
end
