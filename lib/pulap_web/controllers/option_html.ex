defmodule PulapWeb.OptionHTML do
  use PulapWeb, :html

  embed_templates "option_html/*"

  @doc """
  Renders an option form.
  """
  attr :changeset, Ecto.Changeset, required: true
  attr :action, :string, required: true
  attr :set, :map, required: true

  def option_form(assigns) do
    ~H"""
    <.simple_form :let={f} for={@changeset} action={@action}>
      <.error :if={@changeset.action}>
        Oops, something went wrong! Please check the errors below.
      </.error>
      <div class="mt-10 space-y-8 bg-white">
        <.input field={f[:key]} type="text" label="Key" />
        <.input field={f[:label]} type="text" label="Label" />
        <.input field={f[:description]} type="textarea" label="Description" />
        <div class="flex justify-center gap-2 mt-8">
          <.back_button navigate={~p"/sets/#{@set}/options"}>Back</.back_button>
          <.button color="green">
            {if @changeset.data.id, do: "Update", else: "Save"}
          </.button>
        </div>
      </div>
    </.simple_form>
    """
  end
end
