defmodule Pulap.Repo.Migrations.CreateSets do
  use Ecto.Migration

  def change do
    create table(:sets, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :short_code, :string, null: false
      add :key, :string, null: false
      add :label, :string, null: false
      add :description, :text
      add :created_by, :uuid
      add :updated_by, :uuid

      timestamps(type: :utc_datetime)
    end

    create unique_index(:sets, [:short_code])
    create unique_index(:sets, [:key])
  end
end
