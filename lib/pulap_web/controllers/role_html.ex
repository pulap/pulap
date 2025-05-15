defmodule PulapWeb.RoleHTML do
  use PulapWeb, :html

  embed_templates "role_html/*"

  @doc """
  Renders a role form.
  """
  attr :changeset, Ecto.Changeset, required: true
  attr :action, :string, required: true

  def role_form(assigns)
end
