defmodule Pulap.Dict do
  @moduledoc """
  The Dict context.
  """

  import Ecto.Query, warn: false
  alias Pulap.Repo

  alias Pulap.Dict.Dictionary
  alias Pulap.Dict.Entry

  @doc """
  Returns the list of dictionaries.

  ## Examples

      iex> list_dictionaries()
      [%Dictionary{}, ...]

  """
  def list_dictionaries do
    Repo.all(Dictionary)
  end

  @doc """
  Gets a single dictionary.

  Raises `Ecto.NoResultsError` if the Dictionary does not exist.

  ## Examples

      iex> get_dictionary!(123)
      %Dictionary{}

      iex> get_dictionary!(456)
      ** (Ecto.NoResultsError)

  """
  def get_dictionary!(id), do: Repo.get!(Dictionary, id)

  @doc """
  Creates a dictionary.

  ## Examples

      iex> create_dictionary(%{field: value})
      {:ok, %Dictionary{}}

      iex> create_dictionary(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_dictionary(attrs \\ %{}) do
    %Dictionary{}
    |> Dictionary.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a dictionary.

  ## Examples

      iex> update_dictionary(dictionary, %{field: new_value})
      {:ok, %Dictionary{}}

      iex> update_dictionary(dictionary, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_dictionary(%Dictionary{} = dictionary, attrs) do
    dictionary
    |> Dictionary.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a dictionary.

  ## Examples

      iex> delete_dictionary(dictionary)
      {:ok, %Dictionary{}}

      iex> delete_dictionary(dictionary)
      {:error, %Ecto.Changeset{}}

  """
  def delete_dictionary(%Dictionary{} = dictionary) do
    Repo.delete(dictionary)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking dictionary changes.

  ## Examples

      iex> change_dictionary(dictionary)
      %Ecto.Changeset{data: %Dictionary{}}

  """
  def change_dictionary(%Dictionary{} = dictionary, attrs \\ %{}) do
    Dictionary.changeset(dictionary, attrs)
  end

  # Entry functions

  def list_dictionary_entries(%Dictionary{} = dictionary) do
    Entry
    |> where([e], e.dictionary_id == ^dictionary.id)
    |> order_by([e], asc: e.order)
    |> Repo.all()
  end

  def get_dictionary_entry!(id), do: Repo.get!(Entry, id)

  def create_dictionary_entry(attrs \\ %{}) do
    %Entry{}
    |> Entry.changeset(attrs)
    |> Repo.insert()
  end

  def update_dictionary_entry(%Entry{} = entry, attrs) do
    entry
    |> Entry.changeset(attrs)
    |> Repo.update()
  end

  def delete_dictionary_entry(%Entry{} = entry) do
    Repo.delete(entry)
  end

  def change_dictionary_entry(%Entry{} = entry, attrs \\ %{}) do
    Entry.changeset(entry, attrs)
  end
end
