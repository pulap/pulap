defmodule Pulap.Geo.Address do
  use Ecto.Schema
  import Ecto.Changeset
  import Pulap.Utils

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "addresses" do
    field :short_code, :string
    field :name, :string
    field :floor, :string
    field :state, :string
    field :number, :string
    field :street, :string
    field :apartment, :string
    field :city, :string
    field :postal_code, :string
    field :country, :string
    field :admin_level_0, :string
    field :admin_level_1, :string
    field :admin_level_2, :string
    field :admin_level_3, :string
    field :admin_level_4, :string
    field :location_lat, :float
    field :location_lng, :float
    field :geohash, :string
    field :created_by, :binary_id
    field :updated_by, :binary_id

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(address, attrs) do
    address
    |> cast(attrs, [
      :name,
      :street,
      :number,
      :floor,
      :apartment,
      :city,
      :state,
      :postal_code,
      :country,
      :admin_level_0,
      :admin_level_1,
      :admin_level_2,
      :admin_level_3,
      :admin_level_4,
      :location_lat,
      :location_lng,
      :geohash,
      :created_by,
      :updated_by
    ])
    |> validate_required([:street, :city, :state, :country])
    |> put_short_code(:short_code)
  end

  def slug(%__MODULE__{} = address) do
    Pulap.Utils.get_slug(address)
  end
end

defimpl Pulap.SlugSource, for: Pulap.Geo.Address do
  def source_for_slug(%Pulap.Geo.Address{street: street, number: number}) do
    [street, number]
    |> Enum.reject(&is_nil/1)
    |> Enum.join(" ")
  end
end
