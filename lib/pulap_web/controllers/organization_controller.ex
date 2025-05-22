defmodule PulapWeb.OrganizationController do
  use PulapWeb, :controller

  plug :put_layout, {PulapWeb.Layouts, :org}

  alias Pulap.Auth
  alias Pulap.Org.Organization
  alias PulapWeb.AuditHelpers

  def index(conn, _params) do
    redirect(conn, to: "/")
  end

  def new(conn, _params) do
    changeset = Auth.change_organization(%Organization{})
    render(conn, :new, changeset: changeset)
  end

  def create(conn, %{"organization" => organization_params}) do
    params = AuditHelpers.maybe_put_created_by(organization_params, conn)

    case Auth.create_organization(params) do
      {:ok, organization} ->
        conn
        |> put_flash(:info, "Organization created successfully.")
        |> redirect(to: ~p"/organizations/#{organization}")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :new, changeset: changeset)
    end
  end

  def show(conn, %{"id" => id}) do
    case Auth.get_organization(id) do
      nil ->
        conn
        |> put_flash(:error, "Organization not found.")
        |> redirect(to: ~p"/organizations/default")

      organization ->
        organization = Pulap.Repo.preload(organization, :owners)
        render(conn, :show, organization: organization)
    end
  end

  def show_single(conn, _params) do
    case Auth.list_organizations() do
      [organization | _] ->
        organization = Pulap.Repo.preload(organization, :owners)
        render(conn, :show, organization: organization)

      [] ->
        conn
        |> put_flash(:error, "No organization found. Please create one first.")
        |> redirect(to: ~p"/organizations/new")
    end
  end

  def edit(conn, %{"id" => id}) do
    organization = Auth.get_organization!(id)
    changeset = Auth.change_organization(organization)
    render(conn, :edit, organization: organization, changeset: changeset)
  end

  def update(conn, %{"id" => id, "organization" => organization_params}) do
    organization = Auth.get_organization!(id)
    params = AuditHelpers.maybe_put_updated_by(organization_params, conn)

    case Auth.update_organization(organization, params) do
      {:ok, organization} ->
        conn
        |> put_flash(:info, "Organization updated successfully.")
        |> redirect(to: ~p"/organizations/#{organization}")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :edit, organization: organization, changeset: changeset)
    end
  end

  def delete(conn, %{"id" => id}) do
    organization = Auth.get_organization!(id)
    {:ok, _organization} = Auth.delete_organization(organization)

    conn
    |> put_flash(:info, "Organization deleted successfully.")
    |> redirect(to: ~p"/organizations/default")
  end
end
