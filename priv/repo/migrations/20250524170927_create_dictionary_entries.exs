defmodule Pulap.Repo.Migrations.CreateDictionaryEntries do
  use Ecto.Migration

  def change do
    create table(:dictionary_entries, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :label, :string, null: false
      add :description, :text
      add :slug, :string, null: false
      add :value, :string, null: false
      add :order, :integer, default: 0
      add :active, :boolean, default: true

      add :dictionary_id, references(:dictionaries, type: :binary_id, on_delete: :delete_all),
        null: false

      timestamps()
    end

    create index(:dictionary_entries, [:dictionary_id])
    create unique_index(:dictionary_entries, [:slug, :dictionary_id])
  end
end
