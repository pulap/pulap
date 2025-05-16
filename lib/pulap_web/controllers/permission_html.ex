defmodule PulapWeb.PermissionHTML do
  use PulapWeb, :html

  embed_templates "permission_html/*"

  @doc """
  Renders a permission form.
  """
  attr :changeset, Ecto.Changeset, required: true
  attr :action, :string, required: true

  def permission_form(assigns)
end
