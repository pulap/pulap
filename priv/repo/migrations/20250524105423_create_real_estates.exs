defmodule Pulap.Repo.Migrations.CreateRealEstates do
  use Ecto.Migration

  def change do
    create table(:real_estates, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :short_code, :string
      add :name, :string
      add :description, :text
      add :category, :string
      add :type, :string
      add :subtype, :string
      add :surface_total, :float
      add :surface_covered, :float
      add :built_year, :integer
      add :lat, :float
      add :lng, :float
      add :alt, :float
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
      add :geohash, :string
      add :created_by, :uuid
      add :updated_by, :uuid

      timestamps(type: :utc_datetime)
    end
  end
end
