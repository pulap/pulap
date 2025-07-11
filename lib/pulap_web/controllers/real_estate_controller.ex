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
    set = Pulap.Set.get_set_by_key!("estate_category") |> Pulap.Repo.preload(:options)
    categories = Enum.map(set.options, &{&1.label, &1.id})
    types = []
    subtypes = []
    selected_category = nil
    selected_type = nil

    render(conn, :new,
      changeset: changeset,
      categories: categories,
      types: types,
      subtypes: subtypes,
      selected_category: selected_category,
      selected_type: selected_type,
      action: "/real-estates",
      __changed__: nil
    )
  end

  def create(conn, %{"real_estate" => real_estate_params}) do
    params = AuditHelpers.maybe_put_created_by(real_estate_params, conn)
    params = Map.put(params, "updated_by", params["created_by"])

    case Estate.create_real_estate(params) do
      {:ok, real_estate} ->
        conn
        |> put_flash(:info, "Real estate created successfully.")
        |> redirect(to: ~p"/real-estates/#{real_estate}")

      {:error, %Ecto.Changeset{} = changeset} ->
        set = Pulap.Set.get_set_by_key!("estate_category") |> Pulap.Repo.preload(:options)
        categories = Enum.map(set.options, &{&1.label, &1.id})
        types = []
        subtypes = []
        selected_category = nil
        selected_type = nil

        render(conn, :new,
          changeset: changeset,
          categories: categories,
          types: types,
          subtypes: subtypes,
          selected_category: selected_category,
          selected_type: selected_type,
          action: "/real-estates"
        )
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
