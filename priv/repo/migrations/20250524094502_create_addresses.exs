defmodule Pulap.Repo.Migrations.CreateAddresses do
  use Ecto.Migration

  def change do
    create table(:addresses, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :short_code, :string
      add :name, :string
      add :street, :string
      add :number, :string
      add :floor, :string
      add :apartment, :string
      add :city, :string
      add :state, :string
      add :postal_code, :string
      add :country, :string
      add :admin_level_0, :string
      add :admin_level_1, :string
      add :admin_level_2, :string
      add :admin_level_3, :string
      add :admin_level_4, :string
      add :location_lat, :float
      add :location_lng, :float
      add :geohash, :string
      add :created_by, :uuid
      add :updated_by, :uuid

      timestamps(type: :utc_datetime)
    end
  end
end
