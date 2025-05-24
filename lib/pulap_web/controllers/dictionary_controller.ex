defmodule PulapWeb.DictionaryController do
  use PulapWeb, :controller

  alias Pulap.Dict
  alias Pulap.Dict.Dictionary

  def index(conn, _params) do
    dictionaries = Dict.list_dictionaries()
    render(conn, :index, dictionaries: dictionaries)
  end

  def new(conn, _params) do
    changeset = Dict.change_dictionary(%Dictionary{})
    render(conn, :new, changeset: changeset)
  end

  def create(conn, %{"dictionary" => dictionary_params}) do
    IO.inspect(dictionary_params, label: "[DEBUG] Incoming dictionary_params")

    case Dict.create_dictionary(dictionary_params) do
      {:ok, dictionary} ->
        IO.puts("[DEBUG] Dictionary created: #{inspect(dictionary)}")

        conn
        |> put_flash(:info, "Dictionary created successfully.")
        |> redirect(to: ~p"/dictionaries/#{dictionary}")

      {:error, %Ecto.Changeset{} = changeset} ->
        IO.inspect(changeset.errors, label: "[DEBUG] Changeset errors on create")
        render(conn, :new, changeset: changeset)
    end
  end

  def show(conn, %{"id" => id}) do
    dictionary = Dict.get_dictionary!(id)
    render(conn, :show, dictionary: dictionary)
  end

  def edit(conn, %{"id" => id}) do
    dictionary = Dict.get_dictionary!(id)
    changeset = Dict.change_dictionary(dictionary)
    render(conn, :edit, dictionary: dictionary, changeset: changeset)
  end

  def update(conn, %{"id" => id, "dictionary" => dictionary_params}) do
    IO.inspect(dictionary_params, label: "[DEBUG] Incoming dictionary_params (update)")
    dictionary = Dict.get_dictionary!(id)

    case Dict.update_dictionary(dictionary, dictionary_params) do
      {:ok, dictionary} ->
        IO.puts("[DEBUG] Dictionary updated: #{inspect(dictionary)}")

        conn
        |> put_flash(:info, "Dictionary updated successfully.")
        |> redirect(to: ~p"/dictionaries/#{dictionary}")

      {:error, %Ecto.Changeset{} = changeset} ->
        IO.inspect(changeset.errors, label: "[DEBUG] Changeset errors on update")
        render(conn, :edit, dictionary: dictionary, changeset: changeset)
    end
  end

  def delete(conn, %{"id" => id}) do
    dictionary = Dict.get_dictionary!(id)
    {:ok, _dictionary} = Dict.delete_dictionary(dictionary)

    conn
    |> put_flash(:info, "Dictionary deleted successfully.")
    |> redirect(to: ~p"/dictionaries")
  end
end
