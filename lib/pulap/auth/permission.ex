defmodule Pulap.Auth.Permission do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "permissions" do
    field :short_code, :string
    field :name, :string
    field :description, :string
    field :created_by, :binary_id
    field :updated_by, :binary_id

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(permission, attrs) do
    permission
    |> cast(attrs, [:name, :description])
    |> put_short_code(:short_code)
    |> validate_required([:name, :description, :short_code])
    |> unique_constraint(:short_code)
  end

  def slug(%__MODULE__{} = permission) do
    Pulap.Utils.get_slug(permission)
  end
end

defimpl Pulap.SlugSource, for: Pulap.Auth.Permission do
  def source_for_slug(%Pulap.Auth.Permission{name: name}) do
    name
  end
end
