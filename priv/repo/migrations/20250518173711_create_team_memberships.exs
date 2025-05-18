defmodule Pulap.Repo.Migrations.CreateTeamMemberships do
  use Ecto.Migration

  def change do
    create table(:team_memberships, primary_key: false) do
      add :id, :binary_id, primary_key: true
      add :team_id, references(:teams, type: :binary_id, on_delete: :delete_all), null: false
      add :user_id, references(:users, type: :binary_id, on_delete: :delete_all), null: false
      add :relation_type, :string, null: false, default: "direct"
      timestamps(type: :utc_datetime)
    end

    create unique_index(:team_memberships, [:team_id, :user_id])
  end
end
