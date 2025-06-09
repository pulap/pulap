defmodule Pulap.Set do
  @moduledoc """
  The Set context.
  """

  import Ecto.Query, warn: false
  alias Pulap.Repo

  alias Pulap.Set.Set
  alias Pulap.Set.Option

  # Set functions

  @doc """
  Returns the list of sets.

  ## Examples

      iex> list_sets()
      [%Set{}, ...]

  """
  def list_sets do
    Repo.all(Set)
  end

  @doc """
  Gets a single set.

  Raises if the Set does not exist.

  ## Examples

      iex> get_set!(123)
      %Set{}

  """
  def get_set!(id), do: Repo.get!(Set, id)

  @doc """
  Creates a set.

  ## Examples

      iex> create_set(%{field: value})
      {:ok, %Set{}}

  """
  def create_set(attrs \\ %{}) do
    %Set{}
    |> Set.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a set.

  ## Examples

      iex> update_set(set, %{field: new_value})
      {:ok, %Set{}}

  """
  def update_set(%Set{} = set, attrs) do
    set
    |> Set.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a set.

  ## Examples

      iex> delete_set(set)
      {:ok, %Set{}}

  """
  def delete_set(%Set{} = set) do
    Repo.delete(set)
  end

  @doc """
  Returns a changeset for tracking set changes.

  ## Examples

      iex> change_set(set)
      %Ecto.Changeset{data: %Set{}}

  """
  def change_set(%Set{} = set, attrs \\ %{}) do
    Set.changeset(set, attrs)
  end

  # Option functions

  @doc """
  Returns the list of options.

  ## Examples

      iex> list_options()
      [%Option{}, ...]

  """
  def list_options do
    Repo.all(Option)
  end

  @doc """
  Gets a single option.

  Raises `Ecto.NoResultsError` if the Option does not exist.

  ## Examples

      iex> get_option!(123)
      %Option{}

      iex> get_option!(456)
      ** (Ecto.NoResultsError)

  """
  def get_option!(id), do: Repo.get!(Option, id)

  @doc """
  Creates an option.

  ## Examples

      iex> create_option(%{field: value})
      {:ok, %Option{}}

      iex> create_option(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_option(attrs \\ %{}) do
    %Option{}
    |> Option.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates an option.

  ## Examples

      iex> update_option(option, %{field: new_value})
      {:ok, %Option{}}

      iex> update_option(option, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_option(%Option{} = option, attrs) do
    option
    |> Option.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes an option.

  ## Examples

      iex> delete_option(option)
      {:ok, %Option{}}

      iex> delete_option(option)
      {:error, %Ecto.Changeset{}}

  """
  def delete_option(%Option{} = option) do
    Repo.delete(option)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking option changes.

  ## Examples

      iex> change_option(option)
      %Ecto.Changeset{data: %Option{}}

  """
  def change_option(%Option{} = option, attrs \\ %{}) do
    Option.changeset(option, attrs)
  end

  # Nested options functions

  def list_nested_options(%Option{} = parent_option) do
    Option
    |> where([o], o.parent_option_id == ^parent_option.id)
    |> order_by([o], asc: o.order)
    |> Repo.all()
  end

  def get_nested_option!(id), do: Repo.get!(Option, id)

  def create_nested_option(attrs \\ %{}) do
    %Option{}
    |> Option.changeset(attrs)
    |> Repo.insert()
  end

  def update_nested_option(%Option{} = option, attrs) do
    option
    |> Option.changeset(attrs)
    |> Repo.update()
  end

  def delete_nested_option(%Option{} = option) do
    Repo.delete(option)
  end

  def change_nested_option(%Option{} = option, attrs \\ %{}) do
    Option.changeset(option, attrs)
  end
end
