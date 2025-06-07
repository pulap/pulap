defmodule Pulap.Estate.RealEstate do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "real_estates" do
    field :name, :string
    field :type, :string
    field :description, :string
    field :short_code, :string
    field :surface_total, :float
    field :surface_covered, :float
    field :built_year, :integer
    field :lat, :float
    field :lng, :float
    field :alt, :float
    field :created_by, Ecto.UUID
    field :updated_by, Ecto.UUID

    belongs_to :address, Pulap.Geo.Address

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(real_estate, attrs) do
    real_estate
    |> cast(attrs, [
      :name,
      :type,
      :description,
      :surface_total,
      :surface_covered,
      :built_year,
      :lat,
      :lng,
      :alt,
      :created_by,
      :updated_by,
      :address_id
    ])
    |> validate_required([
      :name,
      :type,
      :description,
      :surface_total,
      :surface_covered,
      :built_year,
      :lat,
      :lng,
      :created_by,
      :updated_by
    ])
    |> put_slug(:short_code)
  end
end
