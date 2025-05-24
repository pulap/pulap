defmodule PulapWeb.EntryController do
  use PulapWeb, :controller

  alias Pulap.Dict
  alias Pulap.Dict.Entry

  def index(conn, %{"dictionary_id" => dictionary_id}) do
    dictionary = Dict.get_dictionary!(dictionary_id)
    entries = Dict.list_dictionary_entries(dictionary)
    render(conn, :index, dictionary: dictionary, entries: entries)
  end

  def new(conn, %{"dictionary_id" => dictionary_id}) do
    dictionary = Dict.get_dictionary!(dictionary_id)
    changeset = Dict.change_dictionary_entry(%Entry{dictionary_id: dictionary.id})
    render(conn, :new, dictionary: dictionary, changeset: changeset)
  end

  def create(conn, %{"dictionary_id" => dictionary_id, "entry" => entry_params}) do
    dictionary = Dict.get_dictionary!(dictionary_id)

    case Dict.create_dictionary_entry(entry_params) do
      {:ok, _entry} ->
        conn
        |> put_flash(:info, "Entry created successfully.")
        |> redirect(to: ~p"/dictionaries/#{dictionary}/entries")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :new, dictionary: dictionary, changeset: changeset)
    end
  end

  def edit(conn, %{"dictionary_id" => dictionary_id, "id" => id}) do
    dictionary = Dict.get_dictionary!(dictionary_id)
    entry = Dict.get_dictionary_entry!(id)
    changeset = Dict.change_dictionary_entry(entry)
    render(conn, :edit, dictionary: dictionary, entry: entry, changeset: changeset)
  end

  def update(conn, %{
        "dictionary_id" => dictionary_id,
        "id" => id,
        "entry" => entry_params
      }) do
    dictionary = Dict.get_dictionary!(dictionary_id)
    entry = Dict.get_dictionary_entry!(id)

    case Dict.update_dictionary_entry(entry, entry_params) do
      {:ok, _entry} ->
        conn
        |> put_flash(:info, "Entry updated successfully.")
        |> redirect(to: ~p"/dictionaries/#{dictionary}/entries")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :edit, dictionary: dictionary, entry: entry, changeset: changeset)
    end
  end

  def delete(conn, %{"dictionary_id" => dictionary_id, "id" => id}) do
    dictionary = Dict.get_dictionary!(dictionary_id)
    entry = Dict.get_dictionary_entry!(id)
    {:ok, _entry} = Dict.delete_dictionary_entry(entry)

    conn
    |> put_flash(:info, "Entry deleted successfully.")
    |> redirect(to: ~p"/dictionaries/#{dictionary}/entries")
  end
end
