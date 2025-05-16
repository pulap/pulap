defmodule PulapWeb.ResourceHTML do
  use PulapWeb, :html

  embed_templates "resource_html/*"

  @doc """
  Renders a resource form.
  """
  attr :changeset, Ecto.Changeset, required: true
  attr :action, :string, required: true

  def resource_form(assigns)
end
