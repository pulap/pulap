defmodule Pulap.Repo.Migrations.AddFullAddressToRealEstates do
  use Ecto.Migration

  def change do
    alter table(:real_estates) do
      add :full_address, :string
    end
  end
end
