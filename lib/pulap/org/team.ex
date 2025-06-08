defmodule Pulap.Org.Team do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "teams" do
    field :short_code, :string
    field :name, :string
    field :description, :string
    field :created_by, :binary_id
    field :updated_by, :binary_id
    belongs_to :organization, Pulap.Org.Organization, type: :binary_id
    has_many :team_memberships, Pulap.Org.TeamMembership
    has_many :members, through: [:team_memberships, :user]

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(team, attrs) do
    team
    |> cast(attrs, [:name, :description, :organization_id])
    |> put_short_code(:short_code)
    |> validate_required([:name, :description, :short_code, :organization_id])
    |> unique_constraint(:short_code)
  end

  def slug(%__MODULE__{} = team) do
    Pulap.Utils.get_slug(team)
  end
end

defimpl Pulap.SlugSource, for: Pulap.Org.Team do
  def source_for_slug(%Pulap.Org.Team{name: name}) do
    name
  end
end
