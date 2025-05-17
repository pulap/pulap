defmodule Pulap.Auth.Permission do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "permissions" do
    field :name, :string
    field :description, :string
    field :slug, :string
    field :created_by, :binary_id
    field :updated_by, :binary_id

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(permission, attrs) do
    permission
    |> cast(attrs, [:slug, :name, :description, :created_by, :updated_by])
    |> put_slug()
    |> validate_required([:slug, :name, :description])
    |> unique_constraint(:slug)
  end

  # Generate a slug from the name and a UUID segment, like in Role
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
