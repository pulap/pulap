defmodule Pulap.Auth.ResourceTest do
  use Pulap.DataCase, async: false
  alias Pulap.Auth.Resource

  describe "changeset/2" do
    @valid_attrs %{
      "name" => "API Key",
      "value" => "12345",
      "description" => "API key resource",
      "kind" => "secret"
    }
    @invalid_attrs %{"name" => "", "value" => "", "description" => "", "kind" => ""}

    test "changeset casts only allowed fields and ignores forbidden ones" do
      attrs =
        Map.merge(@valid_attrs, %{
          "created_by" => "should-not-cast",
          "updated_by" => "should-not-cast"
        })

      changeset = Resource.changeset(%Resource{}, attrs)
      assert changeset.changes.name == "API Key"
      assert changeset.changes.value == "12345"
      assert changeset.changes.description == "API key resource"
      assert changeset.changes.kind == "secret"
      assert is_binary(changeset.changes.slug)
      refute Map.has_key?(changeset.changes, :created_by)
      refute Map.has_key?(changeset.changes, :updated_by)
      assert changeset.valid?
    end

    test "validates required fields" do
      changeset = Resource.changeset(%Resource{}, @invalid_attrs)
      refute changeset.valid?

      assert %{
               name: ["can't be blank"],
               value: ["can't be blank"],
               description: ["can't be blank"],
               kind: ["can't be blank"],
               slug: ["can't be blank"]
             } = errors_on(changeset)
    end

    test "generates a slug with the expected format" do
      changeset = Resource.changeset(%Resource{}, @valid_attrs)
      assert changeset.valid?
      slug = Ecto.Changeset.get_field(changeset, :slug)
      assert is_binary(slug)
      assert slug =~ ~r/^api-key-[a-z0-9]+$/
      assert String.length(slug) > 6
    end
  end
end
