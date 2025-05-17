defmodule Pulap.Repo.Migrations.CreateOrganizations do
  use Ecto.Migration

  def change do
    create table(:organizations, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :slug, :string
      add :name, :string
      add :description, :string
      add :created_by, :uuid
      add :updated_by, :uuid
      add :owner_id, references(:users, type: :uuid, on_delete: :nothing)
      timestamps(type: :utc_datetime)
    end

    create unique_index(:organizations, [:slug])
    create index(:organizations, [:owner_id])
  end
end
