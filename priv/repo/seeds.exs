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

email = "superadmin@example.com"

unless Accounts.get_user_by_email(email) do
  # 12-character password
  {:ok, _user} =
    Accounts.register_user(%{
      email: email,
      password: "password1234",
      password_confirmation: "password1234"
    })

  IO.puts("User created: #{email} / password1234")
end
