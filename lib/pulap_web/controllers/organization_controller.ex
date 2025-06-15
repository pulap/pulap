defmodule PulapWeb.OrganizationController do
  use PulapWeb, :controller

  plug :put_layout, {PulapWeb.Layouts, :org}

  alias Pulap.Auth
  alias PulapWeb.AuditHelpers

  def index(conn, _params) do
    redirect(conn, to: ~p"/organizations/default")
  end

  # def new(conn, _params) do
  #   changeset = Auth.change_organization(%Organization{})
  #   render(conn, :new, changeset: changeset)
  # end

  # def create(conn, %{"organization" => organization_params}) do
  #   params = AuditHelpers.maybe_put_created_by(organization_params, conn)
  #
  #   case Auth.create_organization(params) do
  #     {:ok, organization} ->
  #       conn
  #       |> put_flash(:info, "Organization created successfully.")
  #       |> redirect(to: ~p"/organizations/#{organization}")
  #
  #     {:error, %Ecto.Changeset{} = changeset} ->
  #       render(conn, :new, changeset: changeset)
  #   end
  # end

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

  def owners(conn, %{"organization_id" => organization_id}) do
    case Auth.get_organization(organization_id) do
      nil ->
        conn
        |> put_flash(:error, "Organization not found.")
        # Or another appropriate path
        |> redirect(to: ~p"/organizations/default")

      organization ->
        organization = Pulap.Repo.preload(organization, :owners)
        all_users = Auth.list_users()

        assigned_owners_ids = Enum.map(organization.owners, & &1.id)

        available_users =
          Enum.reject(all_users, fn user ->
            user.id in assigned_owners_ids
          end)

        render(conn, :owners,
          organization: organization,
          assigned_owners: organization.owners,
          available_users: available_users
        )
    end
  end

  def assign_owner(conn, %{"organization_id" => organization_id, "user_id" => user_id}) do
    organization = Auth.get_organization!(organization_id)
    user = Auth.get_user!(user_id)

    case Auth.add_organization_owner(organization, user) do
      {:ok, _organization} ->
        conn
        |> put_flash(:info, "Owner assigned successfully.")
        |> redirect(to: ~p"/organizations/#{organization}/owners")

      # Re-added specific handling for :already_owner
      {:error, :already_owner} ->
        conn
        |> put_flash(:error, "Failed to assign owner: This user is already an owner.")
        |> redirect(to: ~p"/organizations/#{organization}/owners")

      {:error, changeset} ->
        conn
        |> put_flash(:error, "Failed to assign owner. Error: #{inspect(changeset.errors)}")
        |> redirect(to: ~p"/organizations/#{organization}/owners")
    end
  end

  def revoke_owner(conn, %{"organization_id" => organization_id, "id" => owner_id}) do
    # Assuming owner_id is the user_id of the owner to be revoked
    organization = Auth.get_organization!(organization_id)
    # owner_id here is actually the user_id
    user = Auth.get_user!(owner_id)

    case Auth.remove_organization_owner(organization, user) do
      {:ok, _organization} ->
        conn
        |> put_flash(:info, "Owner revoked successfully.")
        |> redirect(to: ~p"/organizations/#{organization}/owners")

      # Handle specific atom error
      {:error, :not_an_owner} ->
        conn
        |> put_flash(:error, "Failed to revoke owner: This user is not an owner.")
        |> redirect(to: ~p"/organizations/#{organization}/owners")

      {:error, changeset} ->
        conn
        |> put_flash(:error, "Failed to revoke owner: #{inspect(changeset.errors)}")
        |> redirect(to: ~p"/organizations/#{organization}/owners")
    end
  end
end
