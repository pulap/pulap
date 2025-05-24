# Script for populating the database. You can run it as:
#
#     mix run priv/repo/seeds.exs
#
# Inside the script, you can read and write to any of your
# repositories directly:
#
#     Pulap.Repo.insert!(%Pulap.SomeSchema{})
#
# We recommend using the bang functions (`insert!`, `update!`
# and so on) as they will fail if something goes wrong.

alias Pulap.Accounts
alias Pulap.Repo
alias Pulap.Accounts.User
alias Pulap.Org.Organization
alias Pulap.Org.Team
alias Pulap.Geo.Address
alias Pulap.Estate.RealEstate
alias Pulap.Dict.Dictionary
alias Pulap.Dict.Entry

random_password = fn length ->
  chars = "abcdefghijklmnopqrstuvwxyz0123456789"
  for _ <- 1..length, into: "", do: <<Enum.random(String.to_charlist(chars))>>
end

email = "superadmin@example.com"
password = random_password.(12)

unless Accounts.get_user_by_email(email) do
  {:ok, _user} =
    Accounts.register_user(%{
      email: email,
      username: "superadmin",
      name: "superadmin",
      password: password,
      password_confirmation: password
    })

  IO.puts("User created: #{email} / #{password}")
end

if Repo.aggregate(User, :count, :id) == 0 do
  %User{}
  |> User.registration_changeset(%{
    email: email,
    username: "superadmin",
    name: "superadmin",
    password: password,
    is_active: true
  })
  |> Repo.insert!()

  IO.puts("User created: #{email} / #{password}")
end

org_slug = "default-org"
org_name = "Default Organization"
org_description = "The default organization for the system."

user = Accounts.get_user_by_email(email)

organization =
  Pulap.Repo.get_by(Organization, slug: org_slug) ||
    %Organization{}
    |> Organization.changeset(%{
      slug: org_slug,
      name: org_name,
      short_description: "The default org for all users.",
      description: org_description,
      created_by: user.id
    })
    |> Pulap.Repo.insert!()

# Add superadmin as owner through the join table
Pulap.Repo.insert!(%Pulap.Org.OrganizationOwner{
  organization_id: organization.id,
  user_id: user.id
})

IO.puts("Organization created: #{org_name} (slug: #{org_slug}) and owned by #{user.email}")

# --- Create sample teams and associate them with the organization ---
team_attrs = [
  %{name: "Alpha Team", description: "Handles alpha projects"},
  %{name: "Beta Team", description: "Handles beta testing"},
  %{name: "Gamma Team", description: "Handles gamma operations"}
]

for attrs <- team_attrs do
  Pulap.Auth.create_team(%{
    name: attrs.name,
    description: attrs.description,
    organization_id: organization.id,
    created_by: user.id,
    updated_by: user.id
  })
end

IO.puts("3 sample teams created and associated with organization: #{org_name}")

# --- Create sample address ---
address_attrs = %{
  name: "Sample Address",
  street: "123 Main Street",
  number: "123",
  city: "New York",
  state: "NY",
  country: "USA",
  postal_code: "10001",
  created_by: user.id,
  updated_by: user.id
}

{:ok, address} = Pulap.Geo.create_address(address_attrs)
IO.puts("Sample address created: #{address.street}, #{address.city}")

# --- Create sample real estate ---
real_estate_attrs = %{
  name: "Sample House",
  type: "house",
  description: "A beautiful sample house for testing",
  surface_total: 200.0,
  surface_covered: 150.0,
  built_year: 2020,
  lat: 40.7128,
  lng: -74.0060,
  alt: 10.0,
  address_id: address.id,
  created_by: user.id,
  updated_by: user.id
}

Pulap.Estate.create_real_estate(real_estate_attrs)

IO.puts("Sample real estate created: #{real_estate_attrs.name}")

# Create dictionaries with their entries
dictionaries = [
  %Dictionary{
    label: "Property Types",
    slug: "property-types",
    description: "Types of real estate properties"
  },
  %Dictionary{
    label: "Property Status",
    slug: "property-status",
    description: "Current status of properties"
  },
  %Dictionary{
    label: "Property Features",
    slug: "property-features",
    description: "Available features for properties"
  }
]

# Insert dictionaries and their entries
Enum.each(dictionaries, fn dictionary ->
  {:ok, dict} = Repo.insert!(dictionary)

  entries =
    case dict.slug do
      "property-types" ->
        [
          %Entry{
            value: "house",
            label: "House",
            position: 1,
            active: true,
            dictionary_id: dict.id
          },
          %Entry{
            value: "apartment",
            label: "Apartment",
            position: 2,
            active: true,
            dictionary_id: dict.id
          },
          %Entry{
            value: "commercial",
            label: "Commercial",
            position: 3,
            active: true,
            dictionary_id: dict.id
          }
        ]

      "property-status" ->
        [
          %Entry{
            value: "available",
            label: "Available",
            position: 1,
            active: true,
            dictionary_id: dict.id
          },
          %Entry{value: "sold", label: "Sold", position: 2, active: true, dictionary_id: dict.id},
          %Entry{
            value: "reserved",
            label: "Reserved",
            position: 3,
            active: true,
            dictionary_id: dict.id
          }
        ]

      "property-features" ->
        [
          %Entry{
            value: "garage",
            label: "Garage",
            position: 1,
            active: true,
            dictionary_id: dict.id
          },
          %Entry{
            value: "pool",
            label: "Swimming Pool",
            position: 2,
            active: true,
            dictionary_id: dict.id
          },
          %Entry{
            value: "garden",
            label: "Garden",
            position: 3,
            active: true,
            dictionary_id: dict.id
          }
        ]
    end

  Enum.each(entries, &Repo.insert!(&1))
end)
