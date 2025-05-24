defmodule PulapWeb.EntryHTML do
  use PulapWeb, :html

  embed_templates "entry_html/*"

  attr :changeset, Ecto.Changeset, required: true
  attr :action, :string, required: true
  slot :form_fields, required: true
  slot :actions, required: true

  def entry_form(assigns) do
    ~H"""
    <.form :let={f} for={@changeset} action={@action} phx-submit="save">
      {render_slot(@form_fields, f)}
      {render_slot(@actions)}
    </.form>
    """
  end
end
