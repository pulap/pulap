defmodule Pulap.Geo.AddressTest do
  use Pulap.DataCase, async: true
  alias Pulap.Geo.Address

  describe "changeset/2" do
    @valid_attrs %{
      "slug" => "ul-miodowa-10",
      "name" => "Kamienica przy Miodowej",
      "street" => "ul. Miodowa",
      "number" => "10",
      "floor" => "2",
      "apartment" => "5",
      "city" => "Warszawa",
      "state" => "mazowieckie",
      "postal_code" => "00-251",
      "country" => "Polska",
      "admin_level_0" => "Polska",
      "admin_level_1" => "Mazowieckie",
      "admin_level_2" => "Warszawa",
      "admin_level_3" => "Śródmieście",
      "admin_level_4" => "Stare Miasto",
      "location_lat" => 52.2476,
      "location_lng" => 21.0122,
      "geohash" => "u3h0vj1g",
      "created_by" => "some-uuid",
      "updated_by" => "some-uuid"
    }

    @invalid_attrs %{
      "slug" => nil,
      "street" => nil,
      "number" => nil
    }

    test "valid changeset for Warsaw address" do
      changeset = Address.changeset(%Address{}, @valid_attrs)
      assert changeset.valid?
      assert changeset.changes.city == "Warszawa"
      assert changeset.changes.admin_level_4 == "Stare Miasto"
    end

    test "invalid changeset when required data is missing" do
      changeset = Address.changeset(%Address{}, @invalid_attrs)
      refute changeset.valid?

      errors = errors_on(changeset)
      assert errors[:street] == ["can't be blank"]
      assert errors[:number] == ["can't be blank"]
    end

    test "ignores fields not allowed in cast" do
      attrs = Map.put(@valid_attrs, "unauthorized", "nope")
      changeset = Address.changeset(%Address{}, attrs)
      refute Map.has_key?(changeset.changes, :unauthorized)
      assert changeset.valid?
    end
  end
end
