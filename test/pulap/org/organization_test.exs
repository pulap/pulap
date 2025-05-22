defmodule Pulap.Org.OrganizationTest do
  use Pulap.DataCase, async: false
  alias Pulap.Org.Organization

  describe "changeset/2" do
    @valid_attrs %{
      "name" => "Acme Corp",
      "short_description" => "A company",
      "description" => "Acme Corporation"
    }
    @invalid_attrs %{"name" => "", "short_description" => "", "description" => ""}

    test "changeset casts only allowed fields and ignores forbidden ones" do
      attrs =
        Map.merge(@valid_attrs, %{
          "created_by" => "should-not-cast",
          "updated_by" => "should-not-cast"
        })

      changeset = Organization.changeset(%Organization{}, attrs)
      assert changeset.changes.name == "Acme Corp"
      assert changeset.changes.short_description == "A company"
      assert changeset.changes.description == "Acme Corporation"
      assert is_binary(changeset.changes.slug)
      refute Map.has_key?(changeset.changes, :created_by)
      refute Map.has_key?(changeset.changes, :updated_by)
      assert changeset.valid?
    end

    test "validates required fields" do
      changeset = Organization.changeset(%Organization{}, @invalid_attrs)
      refute changeset.valid?
      assert %{name: ["can't be blank"]} = errors_on(changeset)
    end

    test "generates a slug with the expected format" do
      changeset = Organization.changeset(%Organization{}, @valid_attrs)
      assert changeset.valid?
      slug = Ecto.Changeset.get_field(changeset, :slug)
      assert is_binary(slug)
      assert slug =~ ~r/^acme-corp-[a-z0-9]+$/
      assert String.length(slug) > 6
    end
  end
end
