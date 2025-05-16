defmodule Pulap.Repo.Migrations.CreateResourcePermissions do
  use Ecto.Migration

  def change do
    create table(:resource_permissions, primary_key: false) do
      add :resource_id, references(:resources, type: :binary_id, on_delete: :delete_all), null: false
      add :permission_id, references(:permissions, type: :binary_id, on_delete: :delete_all), null: false
      add :created_by, :binary_id
      add :updated_by, :binary_id
    end

    create unique_index(:resource_permissions, [:resource_id, :permission_id])
  end
end
