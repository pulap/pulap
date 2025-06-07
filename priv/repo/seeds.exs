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

credentials_file = Path.expand("../seeds.out", __DIR__)
File.write!(credentials_file, "", [:write])

write_credentials = fn email, password ->
  IO.puts("User created: #{email} / #{password}")
  File.write!(credentials_file, "#{email} / #{password}\n", [:append])
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

  write_credentials.(email, password)
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

  write_credentials.(email, password)
end

org_code = "def456"
org_name = "Default Organization"
org_description = "The default organization for the system."

user = Accounts.get_user_by_email(email)

organization =
  Pulap.Repo.get_by(Organization, short_code: org_code) ||
    %Organization{}
    |> Organization.changeset(%{
      short_code: org_code,
      name: org_name,
      short_description: "The default org for all users.",
      description: org_description,
      created_by: user.id
    })
    |> Pulap.Repo.insert!()

Pulap.Repo.insert!(%Pulap.Org.OrganizationOwner{
  organization_id: organization.id,
  user_id: user.id
})

IO.puts("Organization created: #{org_name} (code: #{org_code}) and owned by #{user.email}")

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

# Create roles
roles = [
  %{
    name: "System Administrator",
    description: "Full system access",
    scope: "global"
  },
  %{
    name: "Team Leader",
    description: "Can manage team members and assignments",
    scope: "team"
  },
  %{
    name: "Team Member",
    description: "Basic team access",
    scope: "team"
  }
]

created_roles =
  for role_attrs <- roles do
    {:ok, role} = Pulap.Auth.create_role(Map.put(role_attrs, :created_by, user.id))
    role
  end

IO.puts("\nRoles created:")

Enum.each(created_roles, fn role ->
  IO.puts("- #{role.name} (#{role.scope})")
end)

# Assign superadmin to global System Administrator role
[system_admin_role | _] = created_roles
Pulap.Auth.assign_role_to_user(user.id, system_admin_role.id)

IO.puts("\nAssigned #{user.email} to System Administrator role")

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

dictionaries = [
  %Dictionary{
    label: "Property Types",
    short_code: "pt123",
    description: "Types of real estate properties"
  },
  %Dictionary{
    label: "Property Status",
    short_code: "ps456",
    description: "Current status of properties"
  },
  %Dictionary{
    label: "Property Features",
    short_code: "pf789",
    description: "Available features for properties"
  }
]

Enum.each(dictionaries, fn dictionary ->
  dict = Repo.insert!(dictionary)

  entries =
    case dict.short_code do
      "pt123" ->
        [
          %{
            value: "house",
            label: "House",
            order: 1,
            active: true,
            dictionary_id: dict.id,
            created_by: user.id,
            updated_by: user.id
          },
          %{
            value: "apartment",
            label: "Apartment",
            order: 2,
            active: true,
            dictionary_id: dict.id,
            created_by: user.id,
            updated_by: user.id
          },
          %{
            value: "commercial",
            label: "Commercial",
            order: 3,
            active: true,
            dictionary_id: dict.id,
            created_by: user.id,
            updated_by: user.id
          }
        ]

      "ps456" ->
        [
          %{
            value: "available",
            label: "Available",
            order: 1,
            active: true,
            dictionary_id: dict.id,
            created_by: user.id,
            updated_by: user.id
          },
          %{
            value: "sold",
            label: "Sold",
            order: 2,
            active: true,
            dictionary_id: dict.id,
            created_by: user.id,
            updated_by: user.id
          },
          %{
            value: "reserved",
            label: "Reserved",
            order: 3,
            active: true,
            dictionary_id: dict.id,
            created_by: user.id,
            updated_by: user.id
          }
        ]

      "pf789" ->
        [
          %{
            value: "garage",
            label: "Garage",
            order: 1,
            active: true,
            dictionary_id: dict.id,
            created_by: user.id,
            updated_by: user.id
          },
          %{
            value: "pool",
            label: "Swimming Pool",
            order: 2,
            active: true,
            dictionary_id: dict.id,
            created_by: user.id,
            updated_by: user.id
          },
          %{
            value: "garden",
            label: "Garden",
            order: 3,
            active: true,
            dictionary_id: dict.id,
            created_by: user.id,
            updated_by: user.id
          }
        ]
    end

  Enum.each(entries, &Pulap.Dict.create_dictionary_entry/1)
end)
