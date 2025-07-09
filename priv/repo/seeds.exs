alias Pulap.Accounts
alias Pulap.Repo
alias Pulap.Accounts.User

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

_org_code = "def456"
_org_name = "Default Organization"
_org_description = "The default organization for the system."

user = Accounts.get_user_by_email(email)

# ─── Roles ─────────────────────────────────────────────────────────────

IO.puts("\nCreating roles...")

roles = [
  {"Superadmin", false},
  {"Org Admin", false},
  {"Broker Manager", false},
  {"Finance", false},
  {"Support", false},
  {"Auditor", false},
  {"Property Owner", false},
  {"Client", false},
  {"Viewer", false},
  {"Team Lead", true},
  {"Agent", true},
  {"Assistant", true}
]

role_map =
  Enum.reduce(roles, %{}, fn {name, contextual}, acc ->
    role =
      Pulap.Repo.get_by(Pulap.Auth.Role, name: name, contextual: contextual) ||
        elem(
          Pulap.Auth.create_role(%{
            id: Ecto.UUID.generate(),
            name: name,
            description: name,
            contextual: contextual,
            status: "enabled",
            created_by: user.id
          }),
          1
        )

    IO.puts("- #{role.name} (#{if role.contextual, do: "contextual", else: "global"})")
    Map.put(acc, {name, contextual}, role)
  end)

# ─── Permissions ───────────────────────────────────────────────────────

IO.puts("\nCreating permissions...")

permissions = [
  {"Manage Users", false},
  {"Assign Roles", false},
  {"Manage Teams", false},
  {"View Organization Data", false},
  {"Manage Organization", false},
  {"Manage Properties", false},
  {"Approve Listings", false},
  {"Manage Transactions", false},
  {"View Financials", false},
  {"Manage Financials", false},
  {"View Audit Logs", false},
  {"Export Data", false},
  {"Send Mass Notifications", false},
  {"Impersonate Users", false},
  {"Access Admin Dashboard", false},
  {"Create Property", true},
  {"Edit Property Data", true},
  {"Edit Property Location", true},
  {"Manage Property Media", true},
  {"Tag Property Media", true},
  {"Assign Property Agent", true},
  {"View Property History", true},
  {"Create Listing", true},
  {"Edit Listing", true},
  {"Publish Listing", true},
  {"Unpublish Listing", true},
  {"Schedule Visit", true},
  {"Record Client Interaction", true},
  {"Close Deal", true},
  {"Generate Team Report", true},
  {"Comment Internal", true},
  {"Moderate Comments", true},
  {"Reassign Property", true},
  {"Manage Visit Agenda", true}
]

perm_map =
  for {name, _ctx} <- permissions, into: %{} do
    {:ok, perm} =
      Pulap.Auth.create_permission(%{
        id: Ecto.UUID.generate(),
        name: name,
        description: "Permission: #{name}",
        created_by: user.id,
        updated_by: user.id
      })

    {name, perm}
  end

# ─── Role ↔ Permission Assignments ─────────────────────────────────────

IO.puts("\nAssigning permissions to roles...")

assign = fn role_name, ctx, permission_names ->
  Enum.each(permission_names, fn perm_name ->
    Pulap.Auth.assign_permission_to_role(
      role_map[{role_name, ctx}].id,
      perm_map[perm_name].id
    )
  end)
end

assign.("Superadmin", false, Map.keys(perm_map))

assign.("Org Admin", false, [
  "Manage Users",
  "Assign Roles",
  "Manage Teams",
  "View Organization Data",
  "Manage Organization"
])

assign.("Broker Manager", false, [
  "Manage Properties",
  "Approve Listings",
  "Manage Transactions",
  "View Organization Data"
])

assign.("Finance", false, [
  "View Financials",
  "Manage Financials"
])

assign.("Support", false, [
  "Export Data",
  "Send Mass Notifications"
])

assign.("Auditor", false, [
  "View Audit Logs"
])

assign.("Viewer", false, [
  "View Organization Data"
])

assign.("Team Lead", true, [
  "Create Property",
  "Edit Property Data",
  "Create Listing",
  "Assign Property Agent",
  "Schedule Visit",
  "Close Deal",
  "Generate Team Report"
])

assign.("Agent", true, [
  "Create Property",
  "Edit Property Data",
  "Edit Listing",
  "Schedule Visit",
  "Record Client Interaction",
  "Close Deal"
])

assign.("Assistant", true, [
  "Manage Property Media",
  "Tag Property Media",
  "Manage Visit Agenda"
])

# ─── Grant All Global Permissions to Superadmin User ───────────────────

IO.puts("\nGranting all global permissions to Superadmin...")

Enum.each(perm_map, fn {_name, perm} ->
  Pulap.Auth.assign_permission_to_user(user.id, perm.id)
end)

# --- Default Organization and Teams ---
IO.puts("\nCreating default organization and teams...")

# The 'user' variable from earlier (superadmin) is used here.
if user do
  # 1. Create Organization
  org_params = %{
    # Updated Name
    name: "Owners for Real Estate Agency",
    short_description: "The primary real estate agency.",
    # description: "Detailed description of the agency.", # Optional, can be added/edited later
    created_by: user.id
  }

  case Pulap.Auth.create_organization(org_params) do
    # Renamed to avoid confusion
    {:ok, org_from_insert} ->
      IO.puts("- Created Organization: #{org_from_insert.name} (ID: #{org_from_insert.id})")

      IO.puts(
        "  NOTE: Superadmin user (#{user.email}) created this organization via created_by field."
      )

      # Preload the :owners association before attempting to modify it
      org = Pulap.Repo.preload(org_from_insert, :owners)

      # Associate the superadmin user as an owner
      if user do
        # Now org.owners will be an empty list []
        org_with_owner_changeset =
          Ecto.Changeset.change(org)
          |> Ecto.Changeset.put_assoc(:owners, [user])

        case Pulap.Repo.update(org_with_owner_changeset) do
          {:ok, updated_org} ->
            IO.puts(
              "  - Successfully associated #{user.email} as an owner to #{updated_org.name}."
            )

          {:error, owner_changeset} ->
            IO.inspect(owner_changeset, label: "Error associating owner to organization")
        end
      else
        IO.puts("  - Superadmin user not found, cannot associate as owner.")
      end

      IO.puts(
        "  The 'created_by' field links the superadmin. For 'sole ownership', ensure application logic or roles reflect this."
      )

      # 2. Create Teams
      teams_to_create = [
        %{name: "Sales Team A", description: "Primary sales team"},
        %{name: "Sales Team B", description: "Secondary sales team"}
      ]

      Enum.each(teams_to_create, fn team_attrs ->
        team_full_attrs =
          Map.merge(team_attrs, %{
            organization_id: org.id,
            created_by: user.id
          })

        case Pulap.Auth.create_team(team_full_attrs) do
          {:ok, team} ->
            IO.puts("  - Created Team: #{team.name} for Org ID: #{org.id}")

          {:error, changeset} ->
            IO.inspect(changeset, label: "Error creating team #{team_attrs.name}")
        end
      end)

    {:error, changeset} ->
      IO.inspect(changeset, label: "Error creating organization")
  end
else
  IO.puts("Superadmin user (variable 'user') not found. Skipping organization and team creation.")
end

IO.puts("\nCreating sample users for roles and teams...")

create_and_assign_role_for_user = fn email_prefix,
                                     full_name,
                                     role_name_string,
                                     role_contextual_flag ->
  user_email = "#{email_prefix}@example.com"

  base_password = email_prefix

  user_password =
    if String.length(base_password) < 12 do
      base_password <> String.duplicate("0", 12 - String.length(base_password))
    else
      base_password
    end

  # Using email prefix as username for simplicity
  user_username = email_prefix

  role_tuple = {role_name_string, role_contextual_flag}

  unless Accounts.get_user_by_email(user_email) do
    case Accounts.register_user(%{
           email: user_email,
           username: user_username,
           name: full_name,
           password: user_password,
           password_confirmation: user_password
         }) do
      {:ok, new_user_for_role} ->
        write_credentials.(user_email, user_password)

        # Ensure role exists in role_map
        case Map.get(role_map, role_tuple) do
          nil ->
            IO.puts(
              "Error: Role '#{role_name_string}' (contextual: #{role_contextual_flag}) not found in role_map. Skipping role assignment for #{user_email}."
            )

          role_to_assign ->
            Pulap.Auth.assign_role_to_user(new_user_for_role.id, role_to_assign.id)
            IO.puts("- Created user: #{user_email} with role: #{role_name_string}")
        end

      {:error, changeset} ->
        IO.inspect(changeset, label: "Error creating user #{user_email}")
    end
  else
    IO.puts("- User #{user_email} already exists, attempting to assign role if not already.")
    # Optionally, assign role if user exists but role might not be assigned.
    # For simplicity in seeds, we'll skip if user exists to avoid complex checks.
  end
end

# Global Role Users
global_role_users_to_create = [
  {"orgadmin", "Org Admin User", "Org Admin", false},
  {"brokermanager", "Broker Manager User", "Broker Manager", false},
  {"financeuser", "Finance User", "Finance", false},
  {"supportuser", "Support User", "Support", false},
  {"auditoruser", "Auditor User", "Auditor", false},
  {"propertyowner", "Property Owner User", "Property Owner", false},
  {"clientuser", "Client User", "Client", false},
  {"vieweruser", "Viewer User", "Viewer", false}
]

Enum.each(global_role_users_to_create, fn {email_p, name_f, role_n, role_c} ->
  create_and_assign_role_for_user.(email_p, name_f, role_n, role_c)
end)

# Team A Users
team_a_users_to_create = [
  {"team_a_lead", "Team A Lead", "Team Lead", true},
  {"team_a_agent1", "Team A Agent 1", "Agent", true},
  {"team_a_agent2", "Team A Agent 2", "Agent", true},
  {"team_a_agent3", "Team A Agent 3", "Agent", true},
  {"team_a_assistant", "Team A Assistant", "Assistant", true}
]

IO.puts("\nCreating Team A users...")

Enum.each(team_a_users_to_create, fn {email_p, name_f, role_n, role_c} ->
  create_and_assign_role_for_user.(email_p, name_f, role_n, role_c)
end)

# Team B Users
team_b_users_to_create = [
  {"team_b_lead", "Team B Lead", "Team Lead", true},
  {"team_b_agent1", "Team B Agent 1", "Agent", true},
  {"team_b_agent2", "Team B Agent 2", "Agent", true},
  {"team_b_agent3", "Team B Agent 3", "Agent", true},
  {"team_b_assistant", "Team B Assistant", "Assistant", true}
]

IO.puts("\nCreating Team B users...")

Enum.each(team_b_users_to_create, fn {email_p, name_f, role_n, role_c} ->
  create_and_assign_role_for_user.(email_p, name_f, role_n, role_c)
end)

IO.puts(
  "\n✅ Seeding completed (Users, Roles, Permissions, Default Organization & Teams, Sample Users)."
)

# --- Real Estate Categories, Types, and Subtypes ---
IO.puts("\nCreating real estate categories, types, and subtypes...")

estate_category_set =
  Pulap.Repo.get_by(Pulap.Set.Set, key: "estate_category") ||
    Pulap.Set.Set.changeset(%Pulap.Set.Set{}, %{
      id: Ecto.UUID.generate(),
      key: "estate_category",
      label: "Estate Category",
      description: "Top-level real estate categories",
      created_by: user.id,
      updated_by: user.id
    })
    |> Pulap.Repo.insert!()

estate_type_set =
  Pulap.Repo.get_by(Pulap.Set.Set, key: "estate_type") ||
    Pulap.Set.Set.changeset(%Pulap.Set.Set{}, %{
      id: Ecto.UUID.generate(),
      key: "estate_type",
      label: "Estate Type",
      description: "Real estate types",
      created_by: user.id,
      updated_by: user.id
    })
    |> Pulap.Repo.insert!()

estate_subtype_set =
  Pulap.Repo.get_by(Pulap.Set.Set, key: "estate_subtype") ||
    Pulap.Set.Set.changeset(%Pulap.Set.Set{}, %{
      id: Ecto.UUID.generate(),
      key: "estate_subtype",
      label: "Estate Subtype",
      description: "Real estate subtypes",
      created_by: user.id,
      updated_by: user.id
    })
    |> Pulap.Repo.insert!()

categories = %{
  "Residential" => %{
    "House" => [
      "Detached house",
      "Semi-detached house",
      "Bungalow",
      "Ranch-style home",
      "Cottage",
      "Chalet",
      "Cabin",
      "Eco-home",
      "Smart home"
    ],
    "Apartment" => [
      "Studio apartment",
      "Basement apartment",
      "Penthouse",
      "Loft",
      "Condominium",
      "Co-op unit"
    ],
    "Multi-unit" => [
      "Duplex",
      "Triplex",
      "Fourplex",
      "Townhouse",
      "Row house"
    ],
    "Mobile/Modular" => [
      "Mobile home",
      "Modular home",
      "Park model home",
      "Tiny house"
    ],
    "Other" => [
      "Floating home",
      "Houseboat",
      "Yurt",
      "Treehouse"
    ]
  },
  "Commercial" => %{
    "Office" => [
      "Office building",
      "Executive suite",
      "Co-working space"
    ],
    "Retail" => [
      "Retail store",
      "Shopping mall",
      "Strip mall",
      "Showroom"
    ],
    "Hospitality" => [
      "Hotel",
      "Motel",
      "Bed and breakfast",
      "Hostel",
      "Boutique hotel"
    ],
    "Food & Beverage" => [
      "Restaurant",
      "Café",
      "Bar",
      "Nightclub"
    ],
    "Medical" => [
      "Medical office",
      "Dental clinic",
      "Veterinary clinic",
      "Wellness center"
    ],
    "Industrial" => [
      "Warehouse",
      "Factory",
      "Cold storage",
      "Distribution center",
      "Workshop"
    ],
    "Special" => [
      "Gym",
      "Studio",
      "Bank",
      "Data center",
      "Car dealership",
      "Funeral home",
      "Religious facility",
      "School",
      "Government building"
    ]
  },
  "Land" => %{
    "Urban" => [
      "Residential lot",
      "Commercial lot",
      "Corner lot",
      "Infill lot"
    ],
    "Rural" => [
      "Agricultural land",
      "Timberland",
      "Grazing land",
      "Undeveloped land"
    ],
    "Waterfront" => [
      "Beachfront lot",
      "Lakefront lot",
      "Riverfront lot"
    ],
    "Special" => [
      "Mountain land",
      "Desert land",
      "Raw land",
      "Improved land"
    ]
  },
  "Agricultural" => %{
    "Farm" => [
      "Crop farm",
      "Mixed farm",
      "Organic farm"
    ],
    "Ranch" => [
      "Cattle ranch",
      "Horse ranch",
      "Working ranch"
    ],
    "Specialty" => [
      "Orchard",
      "Vineyard",
      "Greenhouse",
      "Fishery",
      "Equestrian estate"
    ]
  },
  "Mixed-use" => %{
    "Live/work" => [
      "Shop-house",
      "Live/work unit"
    ],
    "Mixed building" => [
      "Residential + Commercial",
      "Retail + Office",
      "Office + Industrial"
    ]
  },
  "Special Purpose" => %{
    "Transportation" => [
      "Marina",
      "Boat slip",
      "Dock",
      "Hangar",
      "Railway property",
      "Parking lot",
      "Garage"
    ],
    "Utility / Infrastructure" => [
      "Power station",
      "Water tower",
      "Wind farm",
      "Solar farm",
      "Telecom tower"
    ],
    "Institutional" => [
      "Military facility",
      "Correctional facility",
      "Religious building",
      "Embassy",
      "Library",
      "Post office"
    ],
    "Recreational" => [
      "Event venue",
      "Conference center",
      "Amusement facility"
    ]
  }
}

Enum.each(categories, fn {category, types} ->
  category_key = String.downcase(String.replace(category, ~r/[^a-zA-Z0-9]/, "_"))

  {:ok, category_opt} =
    Pulap.Set.Option.changeset(%Pulap.Set.Option{}, %{
      id: Ecto.UUID.generate(),
      set_id: estate_category_set.id,
      key: category_key,
      label: category,
      value: category,
      parent_id: nil,
      created_by: user.id,
      updated_by: user.id
    })
    |> Pulap.Repo.insert()

  Enum.each(types, fn {type, subtypes} ->
    type_key = category_key <> "_" <> String.downcase(String.replace(type, ~r/[^a-zA-Z0-9]/, "_"))
    type_value = String.downcase(String.replace(type, ~r/[^a-zA-Z0-9]/, "_"))
    type_short_code = "estate-" <> type_value
    type_label = type
    type_option_key = "estate_type_" <> category_key <> "_" <> type_value

    type_descriptions = %{
      "House" =>
        "Standalone residential property, usually with private land and no shared walls.",
      "Apartment" => "A self-contained housing unit in a building with other such units.",
      "Multi-unit" => "A building divided into multiple separate housing units.",
      "Mobile/Modular" => "Homes that are manufactured off-site and moved to their location.",
      "Other" => "Other residential property types not covered above.",
      "Office" => "Commercial property used for business and administrative work.",
      "Retail" => "Property used for selling goods or services directly to consumers.",
      "Hospitality" => "Property for lodging, food, and guest services.",
      "Food & Beverage" => "Property for restaurants, cafes, bars, and similar uses.",
      "Medical" => "Property for healthcare services, clinics, and medical offices.",
      "Industrial" => "Property for manufacturing, storage, or distribution.",
      "Special" => "Specialized commercial property types.",
      "Urban" => "Land located within a city or town, often developed.",
      "Rural" => "Land located outside cities, often used for agriculture or open space.",
      "Waterfront" => "Land adjacent to a body of water.",
      "Farm" => "Land used for growing crops or raising animals.",
      "Ranch" => "Large land area for raising grazing livestock.",
      "Specialty" => "Specialized agricultural property types.",
      "Live/work" => "Mixed-use property for both living and working.",
      "Mixed building" => "Buildings with multiple uses, such as residential and commercial.",
      "Transportation" => "Property used for transport infrastructure.",
      "Utility / Infrastructure" => "Property for utilities and infrastructure.",
      "Institutional" => "Property for public or institutional use.",
      "Recreational" => "Property for recreation and events."
    }

    type_description = Map.get(type_descriptions, type_label, type_label)
    type_order = Enum.find_index(Map.keys(types), fn t -> t == type end) + 1

    {:ok, type_opt} =
      Pulap.Set.Option.changeset(%Pulap.Set.Option{}, %{
        id: Ecto.UUID.generate(),
        set_id: estate_type_set.id,
        short_code: type_short_code,
        key: type_option_key,
        label: type_label,
        description: type_description,
        value: type_value,
        order: type_order,
        active: true,
        parent_id: category_opt.id,
        created_by: user.id,
        updated_by: user.id
      })
      |> Pulap.Repo.insert()

    Enum.each(subtypes, fn subtype ->
      subtype_key =
        type_key <> "_" <> String.downcase(String.replace(subtype, ~r/[^a-zA-Z0-9]/, "_"))

      Pulap.Set.Option.changeset(%Pulap.Set.Option{}, %{
        id: Ecto.UUID.generate(),
        set_id: estate_subtype_set.id,
        key: subtype_key,
        label: subtype,
        value: subtype,
        parent_id: type_opt.id,
        created_by: user.id,
        updated_by: user.id
      })
      |> Pulap.Repo.insert()
    end)
  end)
end)
