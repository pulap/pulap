defmodule Pulap.Estate do
  @moduledoc """
  The Estate context.
  """

  import Ecto.Query, warn: false
  alias Pulap.Repo

  alias Pulap.Estate.RealEstate

  @doc """
  Returns the list of real estates.

  ## Examples

      iex> list_real_estates()
      [%RealEstate{}, ...]

  """
  def list_real_estates do
    RealEstate
    |> Repo.all()
    |> Repo.preload(:address)
  end

  @doc """
  Gets a single real estate.

  Raises `Ecto.NoResultsError` if the Real Estate does not exist.

  ## Examples

      iex> get_real_estate!(123)
      %RealEstate{}

      iex> get_real_estate!(456)
      ** (Ecto.NoResultsError)

  """
  def get_real_estate!(id) do
    RealEstate
    |> Repo.get!(id)
    |> Repo.preload(:address)
  end

  @doc """
  Creates a real estate.

  ## Examples

      iex> create_real_estate(%{field: value})
      {:ok, %RealEstate{}}

      iex> create_real_estate(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_real_estate(attrs \\ %{}) do
    %RealEstate{}
    |> RealEstate.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a real estate.

  ## Examples

      iex> update_real_estate(real_estate, %{field: new_value})
      {:ok, %RealEstate{}}

      iex> update_real_estate(real_estate, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_real_estate(%RealEstate{} = real_estate, attrs) do
    real_estate
    |> RealEstate.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a real estate.

  ## Examples

      iex> delete_real_estate(real_estate)
      {:ok, %RealEstate{}}

      iex> delete_real_estate(real_estate)
      {:error, %Ecto.Changeset{}}

  """
  def delete_real_estate(%RealEstate{} = real_estate) do
    Repo.delete(real_estate)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking real estate changes.

  ## Examples

      iex> change_real_estate(real_estate)
      %Ecto.Changeset{data: %RealEstate{}}

  """
  def change_real_estate(%RealEstate{} = real_estate, attrs \\ %{}) do
    RealEstate.changeset(real_estate, attrs)
  end
end
