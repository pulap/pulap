defmodule Pulap.OrgTest do
  use Pulap.DataCase

  alias Pulap.Org

  describe "organizations" do
    alias Pulap.Org.Organization

    import Pulap.OrgFixtures

    @invalid_attrs %{
      name: nil,
      description: nil,
      slug: nil,
      short_description: nil,
      created_by: nil,
      updated_by: nil
    }

    test "list_organizations/0 returns all organizations" do
      organization = organization_fixture()
      assert Org.list_organizations() == [organization]
    end

    test "get_organization!/1 returns the organization with given id" do
      organization = organization_fixture()
      assert Org.get_organization!(organization.id) == organization
    end

    test "create_organization/1 with valid data creates a organization" do
      valid_attrs = %{
        name: "some name",
        description: "some description",
        slug: "some slug",
        short_description: "some short_description",
        created_by: "some created_by",
        updated_by: "some updated_by"
      }

      assert {:ok, %Organization{} = organization} = Org.create_organization(valid_attrs)
      assert organization.name == "some name"
      # assert organization.slug == "some slug"
      assert organization.description == "some description"
    end

    test "create_organization/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Org.create_organization(@invalid_attrs)
    end

    test "update_organization/2 with valid data updates the organization" do
      organization = organization_fixture()

      update_attrs = %{
        name: "some updated name",
        description: "some updated description",
        slug: "some updated slug",
        short_description: "some updated short_description",
        created_by: "some updated created_by",
        updated_by: "some updated updated_by"
      }

      assert {:ok, %Organization{} = organization} =
               Org.update_organization(organization, update_attrs)

      assert organization.name == "some updated name"
      assert organization.description == "some updated description"

      # assert organization.slug == "some updated slug" # slug is auto-generated, not directly updatable
      assert organization.short_description == "some updated short_description"
      # assert organization.created_by == "some updated created_by" # not cast or updatable
      # assert organization.updated_by == "some updated updated_by" # not cast or updatable
    end

    test "update_organization/2 with invalid data returns error changeset" do
      organization = organization_fixture()
      assert {:error, %Ecto.Changeset{}} = Org.update_organization(organization, @invalid_attrs)
      assert organization == Org.get_organization!(organization.id)
    end

    test "delete_organization/1 deletes the organization" do
      organization = organization_fixture()
      assert {:ok, %Organization{}} = Org.delete_organization(organization)
      assert_raise Ecto.NoResultsError, fn -> Org.get_organization!(organization.id) end
    end

    test "change_organization/1 returns a organization changeset" do
      organization = organization_fixture()
      assert %Ecto.Changeset{} = Org.change_organization(organization)
    end
  end
end
