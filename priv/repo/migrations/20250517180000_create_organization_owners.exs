defmodule Pulap.Repo.Migrations.CreateOrganizationOwners do
  use Ecto.Migration

  def change do
    create table(:organization_owners, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :organization_id, references(:organizations, type: :binary_id, on_delete: :delete_all), null: false
      add :user_id, references(:users, type: :binary_id, on_delete: :delete_all), null: false
      timestamps()
    end
    create unique_index(:organization_owners, [:organization_id, :user_id])
  end
end
