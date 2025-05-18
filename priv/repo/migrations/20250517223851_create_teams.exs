defmodule Pulap.Repo.Migrations.CreateTeams do
  use Ecto.Migration

  def change do
    create table(:teams, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :organization_id, references(:organizations, type: :uuid, on_delete: :delete_all), null: false
      add :slug, :string, null: false
      add :name, :string, null: false
      add :short_description, :string
      add :description, :string
      # add :kind, :string -- removed for consistency
      add :created_by, :uuid
      add :updated_by, :uuid
      timestamps(type: :utc_datetime)
    end

    create unique_index(:teams, [:slug])
  end
end
