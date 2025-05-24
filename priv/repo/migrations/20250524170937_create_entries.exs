defmodule Pulap.Repo.Migrations.CreateEntries do
  use Ecto.Migration

  def change do
    create table(:entries, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :value, :string
      add :label, :string
      add :position, :integer
      add :active, :boolean, default: false, null: false
      add :created_by, :uuid
      add :updated_by, :uuid
      add :dictionary_id, references(:dictionaries, on_delete: :nothing, type: :binary_id)

      timestamps(type: :utc_datetime)
    end

    create index(:entries, [:dictionary_id])
  end
end
