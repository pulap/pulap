defmodule Pulap.Estate.RealEstate do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "real_estates" do
    field :short_code, :string
    field :name, :string
    field :category, :string
    field :type, :string
    field :subtype, :string
    field :description, :string
    field :surface_total, :float
    field :surface_covered, :float
    field :built_year, :integer
    field :lat, :float
    field :lng, :float
    field :alt, :float
    field :street, :string
    field :number, :string
    field :floor, :string
    field :apartment, :string
    field :postal_code, :string
    field :admin_level_0, :string
    field :admin_level_1, :string
    field :admin_level_2, :string
    field :admin_level_3, :string
    field :admin_level_4, :string
    field :geohash, :string
    field :created_by, Ecto.UUID
    field :updated_by, Ecto.UUID

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(real_estate, attrs) do
    real_estate
    |> cast(attrs, [
      :name,
      :category,
      :type,
      :subtype,
      :description,
      :surface_total,
      :surface_covered,
      :built_year,
      :lat,
      :lng,
      :alt,
      :street,
      :number,
      :floor,
      :apartment,
      :postal_code,
      :admin_level_0,
      :admin_level_1,
      :admin_level_2,
      :admin_level_3,
      :admin_level_4,
      :geohash,
      :created_by,
      :updated_by
    ])
    |> validate_required([
      :name,
      :category,
      :type,
      :subtype,
      :description,
      :surface_total,
      :surface_covered,
      :built_year
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
