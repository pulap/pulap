defmodule Pulap.Auth.Role do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "roles" do
    field :name, :string
    field :description, :string
    field :status, :string
    field :slug, :string
    field :created_by, :binary_id
    field :updated_by, :binary_id

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(role, attrs) do
    role
    |> cast(attrs, [:name, :description])
    |> put_status_default()
    |> validate_required([:name, :description])
    |> put_slug()
    |> unique_constraint(:slug)
  end

  defp put_status_default(changeset) do
    case get_field(changeset, :status) do
      nil -> put_change(changeset, :status, "active")
      _ -> changeset
    end
  end
end
