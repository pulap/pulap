defmodule PulapWeb.ResourceController do
  use PulapWeb, :controller

  alias Pulap.Auth
  alias Pulap.Auth.Resource

  def index(conn, _params) do
    resources = Auth.list_resources()
    render(conn, :index, resources: resources)
  end

  def new(conn, _params) do
    changeset = Auth.change_resource(%Resource{})
    render(conn, :new, changeset: changeset)
  end

  def create(conn, %{"resource" => resource_params}) do
    case Auth.create_resource(resource_params) do
      {:ok, resource} ->
        conn
        |> put_flash(:info, "Resource created successfully.")
        |> redirect(to: ~p"/resources/#{resource}")

      {:error, %Ecto.Changeset{} = changeset} ->
        IO.inspect(changeset.errors, label: "[DEBUG] Resource creation failed")
        render(conn, :new, changeset: changeset)
    end
  end

  def show(conn, %{"id" => id}) do
    resource = Auth.get_resource!(id)
    render(conn, :show, resource: resource)
  end

  def edit(conn, %{"id" => id}) do
    resource = Auth.get_resource!(id)
    changeset = Auth.change_resource(resource)
    render(conn, :edit, resource: resource, changeset: changeset)
  end

  def update(conn, %{"id" => id, "resource" => resource_params}) do
    resource = Auth.get_resource!(id)

    case Auth.update_resource(resource, resource_params) do
      {:ok, resource} ->
        conn
        |> put_flash(:info, "Resource updated successfully.")
        |> redirect(to: ~p"/resources/#{resource}")

      {:error, %Ecto.Changeset{} = changeset} ->
        IO.inspect(changeset.errors, label: "[DEBUG] Resource update failed")
        render(conn, :edit, resource: resource, changeset: changeset)
    end
  end

  def delete(conn, %{"id" => id}) do
    resource = Auth.get_resource!(id)
    {:ok, _resource} = Auth.delete_resource(resource)

    conn
    |> put_flash(:info, "Resource deleted successfully.")
    |> redirect(to: ~p"/resources")
  end

  def permissions(conn, %{"id" => resource_id}) do
    resource = Auth.get_resource!(resource_id)
    permissions = Auth.get_permissions_with_assignment_status_for_resource(resource_id)
    {assigned, unassigned} = Enum.split_with(permissions, & &1.assigned)
    render(conn, "permissions.html", resource: resource, assigned_permissions: assigned, unassigned_permissions: unassigned)
  end

  def assign_permission(conn, %{"id" => resource_id, "permission_id" => permission_id}) do
    case Auth.assign_permission_to_resource(resource_id, permission_id) do
      {:ok, _} ->
        conn
        |> put_flash(:info, "Permission assigned successfully.")
        |> redirect(to: ~p"/resources/#{resource_id}/permissions")
      {:error, :already_assigned} ->
        conn
        |> put_flash(:error, "Permission was already assigned.")
        |> redirect(to: ~p"/resources/#{resource_id}/permissions")
      {:error, _} ->
        conn
        |> put_flash(:error, "Could not assign permission.")
        |> redirect(to: ~p"/resources/#{resource_id}/permissions")
    end
  end

  def revoke_permission(conn, %{"id" => resource_id, "permission_id" => permission_id}) do
    case Auth.revoke_permission_from_resource(resource_id, permission_id) do
      {:ok, _} ->
        conn
        |> put_flash(:info, "Permission revoked successfully.")
        |> redirect(to: ~p"/resources/#{resource_id}/permissions")
      {:error, _} ->
        conn
        |> put_flash(:error, "Could not revoke permission.")
        |> redirect(to: ~p"/resources/#{resource_id}/permissions")
    end
  end
end
