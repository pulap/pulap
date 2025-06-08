defmodule Pulap.Repo.Migrations.CreateDictionaries do
  use Ecto.Migration

  def change do
    create table(:dictionaries, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :short_code, :string
      add :label, :string
      add :description, :text
      add :created_by, :uuid
      add :updated_by, :uuid

      timestamps(type: :utc_datetime)
    end

    create unique_index(:dictionaries, [:short_code])
  end
end
