defmodule Pulap.Auth.Team do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "teams" do
    field :name, :string
    field :description, :string
    field :kind, :string
    field :slug, :string
    field :created_by, :binary_id
    field :updated_by, :binary_id
    belongs_to :organization, Pulap.Auth.Organization, type: :binary_id

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(team, attrs) do
    team
    |> cast(attrs, [:name, :description, :kind, :created_by, :updated_by, :organization_id])
    |> put_slug()
    |> validate_required([:name, :description, :kind, :created_by, :updated_by, :organization_id, :slug])
  end

  defp put_slug(changeset) do
    if name = get_field(changeset, :name) do
      uuid = Ecto.UUID.generate()
      last_segment = uuid |> String.split("-") |> List.last()
      slug =
        name
        |> String.downcase()
        |> String.replace(~r/[^a-z0-9]+/u, "-")
        |> String.trim("-")
        |> Kernel.<>("-" <> last_segment)
      put_change(changeset, :slug, slug)
    else
      changeset
    end
  end
end
