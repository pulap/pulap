defmodule Pulap.Repo.Migrations.CreateRolePermissions do
  use Ecto.Migration

  def change do
    create table(:role_permissions, primary_key: false) do
      add :role_id, references(:roles, type: :binary_id, on_delete: :delete_all), null: false
      add :permission_id, references(:permissions, type: :binary_id, on_delete: :delete_all), null: false
    end

    create unique_index(:role_permissions, [:role_id, :permission_id])
  end
end