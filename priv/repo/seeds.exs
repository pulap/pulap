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
alias Pulap.Auth.Organization

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
