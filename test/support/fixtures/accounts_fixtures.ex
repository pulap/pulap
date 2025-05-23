defmodule Pulap.AccountsFixtures do
  @moduledoc """
  This module defines test helpers for creating
  entities via the `Pulap.Accounts` context.
  """

  def unique_user_email, do: "user#{System.unique_integer()}@example.com"
  def valid_user_password, do: "hello world!"

  def valid_user_attributes(attrs \\ %{}) do
    Enum.into(attrs, %{
      email: unique_user_email(),
      password: valid_user_password(),
      username: "user_#{System.unique_integer()}",
      name: "Test User"
    })
  end

  def user_fixture(attrs \\ %{}) do
    {:ok, user} =
      attrs
      |> valid_user_attributes()
      |> Pulap.Accounts.register_user()

    user
  end

  def extract_user_token(fun) do
    {:ok, captured_email} = fun.(&"[TOKEN]#{&1}[TOKEN]")
    [_, token | _] = String.split(captured_email.text_body, "[TOKEN]")
    token
  end

  @doc """
  Generate a user.
  """
  def user_fixture_create(attrs \\ %{}) do
    {:ok, user} =
      attrs
      |> Enum.into(%{
        confirmed_at: ~U[2025-05-15 18:05:00Z],
        email: "some email",
        hashed_password: "some hashed_password"
      })
      |> Pulap.Accounts.create_user()

    user
  end
end
