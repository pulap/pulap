defmodule Pulap.Repo do
  use Ecto.Repo,
    otp_app: :pulap,
    adapter: Ecto.Adapters.SQLite3
end
