defmodule Pulap.Auth.Resource do
  use Ecto.Schema
  import Ecto.Changeset

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
    |> cast(attrs, [:name, :description, :kind, :value, :created_by, :updated_by])
    |> put_slug()
    |> validate_required([:name, :description, :kind, :value])
    |> unique_constraint(:slug)
  end

  # Generate a slug from the name and a UUID segment, like in Role/Permission
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
      change(changeset, slug: slug)
    else
      changeset
    end
  end
end
