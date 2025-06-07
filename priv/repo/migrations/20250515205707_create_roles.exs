defmodule Pulap.Repo.Migrations.CreateRoles do
  use Ecto.Migration

  def change do
    create table(:roles, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :short_code, :string
      add :name, :string, null: false
      add :description, :text
      add :scope, :string
      add :organization_id, references(:organizations, type: :binary_id)
      add :status, :string, default: "active"

      timestamps(type: :utc_datetime)
    end

    create unique_index(:roles, [:name, :organization_id, :scope])
  end
end
