defmodule PulapWeb.OptionController do
  use PulapWeb, :controller

  alias Pulap.Set
  alias Pulap.Set.Option

  def index(conn, %{"set_id" => set_id}) do
    set = Set.get_set!(set_id)
    options = Set.list_options()
    render(conn, :index, options: options, set: set)
  end

  def new(conn, %{"set_id" => set_id}) do
    set = Set.get_set!(set_id)
    changeset = Set.change_option(%Option{set_id: set_id})
    render(conn, :new, changeset: changeset, set: set)
  end

  def create(conn, %{"set_id" => set_id, "option" => option_params}) do
    set = Set.get_set!(set_id)

    case Set.create_option(Map.put(option_params, "set_id", set_id)) do
      {:ok, option} ->
        conn
        |> put_flash(:info, "Option created successfully.")
        |> redirect(to: ~p"/sets/#{set_id}/options/#{option}")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :new, changeset: changeset, set: set)
    end
  end

  def show(conn, %{"set_id" => set_id, "id" => id}) do
    set = Set.get_set!(set_id)
    option = Set.get_option!(id)
    render(conn, :show, option: option, set: set)
  end

  def edit(conn, %{"set_id" => set_id, "id" => id}) do
    set = Set.get_set!(set_id)
    option = Set.get_option!(id)
    changeset = Set.change_option(option)
    render(conn, :edit, option: option, changeset: changeset, set: set)
  end

  def update(conn, %{"set_id" => set_id, "id" => id, "option" => option_params}) do
    set = Set.get_set!(set_id)
    option = Set.get_option!(id)

    case Set.update_option(option, option_params) do
      {:ok, option} ->
        conn
        |> put_flash(:info, "Option updated successfully.")
        |> redirect(to: ~p"/sets/#{set_id}/options/#{option}")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :edit, option: option, changeset: changeset, set: set)
    end
  end

  def delete(conn, %{"set_id" => set_id, "id" => id}) do
    _set = Set.get_set!(set_id)
    option = Set.get_option!(id)
    {:ok, _option} = Set.delete_option(option)

    conn
    |> put_flash(:info, "Option deleted successfully.")
    |> redirect(to: ~p"/sets/#{set_id}/options")
  end
end
