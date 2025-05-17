defmodule Pulap.Org.OrganizationOwner do
  use Ecto.Schema

  @primary_key {:id, :binary_id, autogenerate: true}
  schema "organization_owners" do
    field :organization_id, :binary_id
    field :user_id, :binary_id
    timestamps()
  end
end
