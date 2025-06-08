defmodule Pulap.Repo.Migrations.CreateUsersAuthTables do
  use Ecto.Migration

  def change do
    create table(:users, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :short_code, :string
      add :name, :string
      add :username, :string
      add :email, :string, null: false, collate: :nocase
      add :hashed_password, :string, null: false
      add :last_login_at, :utc_datetime
      add :last_login_ip, :string
      add :is_active, :boolean, default: true, null: false
      add :confirmed_at, :utc_datetime
      add :created_by, :binary_id
      add :updated_by, :binary_id

      timestamps(type: :utc_datetime)
    end

    create unique_index(:users, [:email])
    create unique_index(:users, [:short_code])

    create table(:users_tokens, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :user_id, references(:users, type: :binary_id, on_delete: :delete_all), null: false
      add :token, :binary, null: false, size: 32
      add :context, :string, null: false
      add :sent_to, :string

      timestamps(type: :utc_datetime, updated_at: false)
    end

    create index(:users_tokens, [:user_id])
    create unique_index(:users_tokens, [:context, :token])
  end
end
