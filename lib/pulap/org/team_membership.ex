defmodule Pulap.Org.TeamMembership do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  schema "team_memberships" do
    belongs_to :team, Pulap.Org.Team, type: :binary_id
    belongs_to :user, Pulap.Accounts.User, type: :binary_id
    field :relation_type, :string, default: "direct"
    timestamps(type: :utc_datetime)
  end

  def changeset(team_membership, attrs) do
    team_membership
    |> cast(attrs, [:team_id, :user_id, :relation_type])
    |> validate_required([:team_id, :user_id, :relation_type])
    |> unique_constraint([:team_id, :user_id])
  end
end
