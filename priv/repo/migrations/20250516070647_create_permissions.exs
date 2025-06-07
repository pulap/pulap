defmodule Pulap.Repo.Migrations.CreatePermissions do
  use Ecto.Migration

  def change do
    create table(:permissions, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :short_code, :string
      add :name, :string
      add :description, :string
      add :created_by, :binary_id
      add :updated_by, :binary_id

      timestamps(type: :utc_datetime)
    end

    create index(:permissions, [:created_by])
    create index(:permissions, [:updated_by])
  end
end
