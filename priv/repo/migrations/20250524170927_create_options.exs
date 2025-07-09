defmodule Pulap.Repo.Migrations.CreateOptions do
  use Ecto.Migration

  def change do
    create table(:options, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :short_code, :string, null: false
      add :key, :string, null: false
      add :label, :string, null: false
      add :description, :text
      add :value, :string, null: false
      add :order, :integer, default: 0
      add :active, :boolean, default: true

      add :set_id, references(:sets, type: :binary_id, on_delete: :delete_all), null: false
      add :parent_id, references(:options, type: :binary_id, on_delete: :nilify_all)

      add :created_by, :uuid
      add :updated_by, :uuid

      timestamps(type: :utc_datetime)
    end

    create index(:options, [:set_id])
    create index(:options, [:parent_id])
    create unique_index(:options, [:short_code, :set_id])
    create unique_index(:options, [:key, :set_id])
  end
end
