defmodule PulapWeb.SetController do
  use PulapWeb, :controller

  alias Pulap.Set
  alias Pulap.Set.Set, as: SetSchema

  def index(conn, _params) do
    sets = Set.list_sets()
    render(conn, :index, sets: sets)
  end

  def new(conn, _params) do
    changeset = Set.change_set(%SetSchema{})
    render(conn, :new, changeset: changeset)
  end

  def create(conn, %{"set" => set_params}) do
    case Set.create_set(set_params) do
      {:ok, set} ->
        conn
        |> put_flash(:info, "Set created successfully.")
        |> redirect(to: ~p"/sets/#{set}")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :new, changeset: changeset)
    end
  end

  def show(conn, %{"id" => id}) do
    set = Set.get_set!(id)
    render(conn, :show, set: set)
  end

  def edit(conn, %{"id" => id}) do
    set = Set.get_set!(id)
    changeset = Set.change_set(set)
    render(conn, :edit, set: set, changeset: changeset)
  end

  def update(conn, %{"id" => id, "set" => set_params}) do
    set = Set.get_set!(id)

    case Set.update_set(set, set_params) do
      {:ok, set} ->
        conn
        |> put_flash(:info, "Set updated successfully.")
        |> redirect(to: ~p"/sets/#{set}")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :edit, set: set, changeset: changeset)
    end
  end

  def delete(conn, %{"id" => id}) do
    set = Set.get_set!(id)
    {:ok, _set} = Set.delete_set(set)

    conn
    |> put_flash(:info, "Set deleted successfully.")
    |> redirect(to: ~p"/sets")
  end
end
