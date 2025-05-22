defmodule Pulap.Auth.RoleTest do
  use Pulap.DataCase, async: false
  alias Pulap.Auth.Role

  describe "changeset/2" do
    @valid_attrs %{"name" => "Admin", "description" => "Admin role"}
    @invalid_attrs %{"name" => "", "description" => ""}

    test "changeset casts only allowed fields and ignores forbidden ones" do
      attrs =
        Map.merge(@valid_attrs, %{
          "created_by" => "should-not-cast",
          "updated_by" => "should-not-cast"
        })

      changeset = Role.changeset(%Role{}, attrs)
      assert changeset.changes.name == "Admin"
      assert changeset.changes.description == "Admin role"
      assert changeset.changes.status == "active"
      assert is_binary(changeset.changes.slug)
      refute Map.has_key?(changeset.changes, :created_by)
      refute Map.has_key?(changeset.changes, :updated_by)
      assert changeset.valid?
    end

    test "validates required fields" do
      changeset = Role.changeset(%Role{}, @invalid_attrs)
      refute changeset.valid?
      assert %{name: ["can't be blank"], description: ["can't be blank"]} = errors_on(changeset)
    end

    test "generates a slug with the expected format" do
      changeset = Role.changeset(%Role{}, @valid_attrs)
      assert changeset.valid?
      slug = Ecto.Changeset.get_field(changeset, :slug)
      assert is_binary(slug)
      assert slug =~ ~r/^admin-[a-z0-9]+$/
      assert String.length(slug) > 6
    end

    # test "enforces unique constraint on slug (integration)" do
    #   role1 = %Role{} |> Role.changeset(@valid_attrs) |> Pulap.Repo.insert!()
    #   changeset2 = Role.changeset(%Role{}, @valid_attrs)
    #   assert {:error, changeset} = Pulap.Repo.insert(changeset2)
    #   assert {:slug, {_, [constraint: :unique, constraint_name: _]}} = hd(changeset.errors)
    #   refute changeset.valid?
    # end
  end
end
