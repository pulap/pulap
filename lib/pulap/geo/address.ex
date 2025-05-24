defmodule Pulap.Geo.Address do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "addresses" do
    field :name, :string
    field :floor, :string
    field :state, :string
    field :number, :string
    field :slug, :string
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
    field :created_by, Ecto.UUID
    field :updated_by, Ecto.UUID

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
      :admin_level_4
    ])
    # |> put_slug()
    # |> put_geolocation()
    # |> put_audit_fields()
    |> validate_required([:name, :street, :number, :city, :state, :postal_code, :country])
  end
end
