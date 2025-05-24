defmodule PulapWeb.RealEstateController do
  use PulapWeb, :controller

  plug :put_layout, html: {PulapWeb.Layouts, :core}

  alias Pulap.Estate
  alias Pulap.Estate.RealEstate
  alias PulapWeb.AuditHelpers

  def index(conn, _params) do
    real_estates = Estate.list_real_estates()
    render(conn, :index, real_estates: real_estates)
  end

  def new(conn, _params) do
    changeset = Estate.change_real_estate(%RealEstate{})
    render(conn, :new, changeset: changeset)
  end

  def create(conn, %{"real_estate" => real_estate_params}) do
    params = AuditHelpers.maybe_put_created_by(real_estate_params, conn)
    params = Map.put(params, "updated_by", params["created_by"])

    # WIP: Temporary solution to handle the required address_id field until the proper address
    # selection/creation flow is implemented in the UI
    params = case Map.get(params, "address_id") do
      nil ->
        {:ok, address} = Pulap.Geo.create_address(%{
          name: "Default Address",
          street: "Default Street",
          number: "1",
          city: "Default City",
          state: "Default State",
          country: "Default Country",
          postal_code: "00000",
          created_by: params["created_by"],
          updated_by: params["updated_by"]
        })
        Map.put(params, "address_id", address.id)
      _ -> params
    end

    case Estate.create_real_estate(params) do
      {:ok, real_estate} ->
        conn
        |> put_flash(:info, "Real estate created successfully.")
        |> redirect(to: ~p"/real-estates/#{real_estate}")

      {:error, %Ecto.Changeset{} = changeset} ->
        IO.inspect(changeset.errors, label: "[DEBUG] Real estate creation failed")
        render(conn, :new, changeset: changeset)
    end
  end

  def show(conn, %{"id" => id}) do
    real_estate = Estate.get_real_estate!(id)
    render(conn, :show, real_estate: real_estate)
  end

  def edit(conn, %{"id" => id}) do
    real_estate = Estate.get_real_estate!(id)
    changeset = Estate.change_real_estate(real_estate)
    render(conn, :edit, real_estate: real_estate, changeset: changeset)
  end

  def update(conn, %{"id" => id, "real_estate" => real_estate_params}) do
    real_estate = Estate.get_real_estate!(id)
    params = AuditHelpers.maybe_put_updated_by(real_estate_params, conn)

    case Estate.update_real_estate(real_estate, params) do
      {:ok, real_estate} ->
        conn
        |> put_flash(:info, "Real estate updated successfully.")
        |> redirect(to: ~p"/real-estates/#{real_estate}")

      {:error, %Ecto.Changeset{} = changeset} ->
        IO.inspect(changeset.errors, label: "[DEBUG] Real estate update failed")
        render(conn, :edit, real_estate: real_estate, changeset: changeset)
    end
  end

  def delete(conn, %{"id" => id}) do
    real_estate = Estate.get_real_estate!(id)
    {:ok, _real_estate} = Estate.delete_real_estate(real_estate)

    conn
    |> put_flash(:info, "Real estate deleted successfully.")
    |> redirect(to: ~p"/real-estates")
  end
end
