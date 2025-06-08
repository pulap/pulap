defmodule Pulap.Auth.Role do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "roles" do
    field :short_code, :string
    field :name, :string
    field :description, :string
    field :status, :string, default: "active"
    field :contextual, :boolean, default: false

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(role, attrs) do
    role
    |> cast(attrs, [:name, :description, :contextual])
    |> put_status_default()
    |> validate_required([:name, :description])
    |> unique_constraint(:name, name: :roles_name_contextual_index)
  end

  defp put_status_default(changeset) do
    case get_field(changeset, :status) do
      nil -> put_change(changeset, :status, "active")
      _ -> changeset
    end
  end
end
