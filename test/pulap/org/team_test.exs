defmodule Pulap.Org.TeamTest do
  use Pulap.DataCase, async: false
  alias Pulap.Org.Team

  describe "changeset/2" do
    @valid_attrs %{"name" => "Dev Team", "description" => "Development team"}
    @invalid_attrs %{"name" => "", "description" => ""}

    test "changeset casts only allowed fields and ignores forbidden ones" do
      attrs =
        Map.merge(@valid_attrs, %{
          "created_by" => "should-not-cast",
          "updated_by" => "should-not-cast"
        })

      changeset = Team.changeset(%Team{}, attrs)
      assert changeset.changes.name == "Dev Team"
      assert changeset.changes.description == "Development team"
      assert is_binary(changeset.changes.slug)
      refute Map.has_key?(changeset.changes, :created_by)
      refute Map.has_key?(changeset.changes, :updated_by)
      assert changeset.valid?
    end

    test "validates required fields" do
      changeset = Team.changeset(%Team{}, @invalid_attrs)
      refute changeset.valid?

      assert %{
               name: ["can't be blank"],
               description: ["can't be blank"],
               slug: ["can't be blank"]
             } = errors_on(changeset)
    end

    test "generates a slug with the expected format" do
      changeset = Team.changeset(%Team{}, @valid_attrs)
      assert changeset.valid?
      slug = Ecto.Changeset.get_field(changeset, :slug)
      assert is_binary(slug)
      assert slug =~ ~r/^dev-team-[a-z0-9]+$/
      assert String.length(slug) > 6
    end

    # Note: Unique constraint test is omitted due to SQLite limitations.
    # If you use Postgres, you can add a test for unique constraint on slug.
  end
end
