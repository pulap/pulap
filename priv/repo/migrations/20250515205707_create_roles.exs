defmodule Pulap.Repo.Migrations.CreateRoles do
  use Ecto.Migration

  def change do
    create table(:roles, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :short_code, :string
      add :name, :string
      add :description, :string
      add :status, :string
      add :created_by, :binary_id
      add :updated_by, :binary_id

      timestamps(type: :utc_datetime)
    end

    create unique_index(:roles, [:short_code])
  end
end
