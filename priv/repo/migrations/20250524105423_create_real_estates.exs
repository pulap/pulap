defmodule Pulap.Repo.Migrations.CreateRealEstates do
  use Ecto.Migration

  def change do
    create table(:real_estates, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :short_code, :string
      add :name, :string
      add :description, :text
      add :type, :string
      add :surface_total, :float
      add :surface_covered, :float
      add :built_year, :integer
      add :lat, :float
      add :lng, :float
      add :alt, :float
      add :created_by, :uuid
      add :updated_by, :uuid
      add :address_id, references(:addresses, on_delete: :nothing, type: :binary_id)

      timestamps(type: :utc_datetime)
    end

    create index(:real_estates, [:address_id])
  end
end
