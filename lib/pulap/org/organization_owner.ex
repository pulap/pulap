defmodule Pulap.Org.OrganizationOwner do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  schema "organization_owners" do
    belongs_to :organization, Pulap.Org.Organization, type: :binary_id
    belongs_to :user, Pulap.Accounts.User, type: :binary_id

    # Enable both inserted_at and updated_at
    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(attrs) do
    %__MODULE__{}
    |> cast(attrs, [:organization_id, :user_id])
    |> validate_required([:organization_id, :user_id])
    |> foreign_key_constraint(:organization_id)
    |> foreign_key_constraint(:user_id)
    |> unique_constraint([:organization_id, :user_id])
  end

  def creation_changeset(organization, user) do
    %__MODULE__{}
    |> cast(%{organization_id: organization.id, user_id: user.id}, [:organization_id, :user_id])
    |> validate_required([:organization_id, :user_id])
    |> foreign_key_constraint(:organization_id)
    |> foreign_key_constraint(:user_id)
    |> unique_constraint([:organization_id, :user_id])
  end
end
