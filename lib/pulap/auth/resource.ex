defmodule Pulap.Auth.Resource do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "resources" do
    field :short_code, :string
    field :name, :string
    field :value, :string
    field :description, :string
    field :kind, :string
    field :created_by, Ecto.UUID
    field :updated_by, Ecto.UUID

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(resource, attrs) do
    resource
    |> cast(attrs, [:name, :value, :description, :kind])
    |> put_short_code(:short_code)
    |> validate_required([:name, :value, :description, :kind, :short_code])
    |> unique_constraint(:short_code)
  end

  def slug(%__MODULE__{} = resource) do
    Pulap.Utils.get_slug(resource)
  end
end

defimpl Pulap.SlugSource, for: Pulap.Auth.Resource do
  def source_for_slug(%Pulap.Auth.Resource{name: name}) do
    name
  end
end
