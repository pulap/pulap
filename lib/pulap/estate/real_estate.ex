defmodule Pulap.Estate.RealEstate do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "real_estates" do
    field :short_code, :string
    field :name, :string
    field :type, :string
    field :description, :string
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
      :address_id,
      :created_by,
      :updated_by
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
    |> put_short_code(:short_code)
  end

  def slug(%__MODULE__{} = real_estate) do
    Pulap.Utils.get_slug(real_estate)
  end
end

defimpl Pulap.SlugSource, for: Pulap.Estate.RealEstate do
  def source_for_slug(%Pulap.Estate.RealEstate{name: name}) do
    name
  end
end
