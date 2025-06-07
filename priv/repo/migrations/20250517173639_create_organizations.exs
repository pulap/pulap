defmodule Pulap.Repo.Migrations.CreateOrganizations do
  use Ecto.Migration

  def change do
    create table(:organizations, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :short_code, :string
      add :name, :string
      add :short_description, :string
      add :description, :string
      add :created_by, :uuid
      add :updated_by, :uuid
      timestamps(type: :utc_datetime)
    end

    create unique_index(:organizations, [:short_code])
  end
end
