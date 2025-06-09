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
  "Manage Users", "Assign Roles", "Manage Teams",
  "View Organization Data", "Manage Organization"
])

assign.("Broker Manager", false, [
  "Manage Properties", "Approve Listings", "Manage Transactions",
  "View Organization Data"
])

assign.("Finance", false, [
  "View Financials", "Manage Financials"
])

assign.("Support", false, [
  "Export Data", "Send Mass Notifications"
])

assign.("Auditor", false, [
  "View Audit Logs"
])

assign.("Viewer", false, [
  "View Organization Data"
])

assign.("Team Lead", true, [
  "Create Property", "Edit Property Data", "Create Listing",
  "Assign Property Agent", "Schedule Visit", "Close Deal",
  "Generate Team Report"
])

assign.("Agent", true, [
  "Create Property", "Edit Property Data", "Edit Listing",
  "Schedule Visit", "Record Client Interaction", "Close Deal"
])

assign.("Assistant", true, [
  "Manage Property Media", "Tag Property Media", "Manage Visit Agenda"
])

# ─── Grant All Global Permissions to Superadmin User ───────────────────

IO.puts("\nGranting all global permissions to Superadmin...")

Enum.each(perm_map, fn {_name, perm} ->
  Pulap.Auth.assign_permission_to_user(user.id, perm.id)
end)

IO.puts("\n✅ Roles, permissions and assignments completed.")