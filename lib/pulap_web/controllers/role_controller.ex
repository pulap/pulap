defmodule PulapWeb.RoleController do
  use PulapWeb, :controller

  plug :put_layout, {PulapWeb.Layouts, :auth}

  alias Pulap.Auth
  alias Pulap.Auth.Role
  alias PulapWeb.AuditHelpers

  def index(conn, _params) do
    roles = Auth.list_roles()
    render(conn, :index, roles: roles)
  end

  def new(conn, _params) do
    changeset = Auth.change_role(%Role{})
    render(conn, :new, changeset: changeset)
  end

  def create(conn, %{"role" => role_params}) do
    params = AuditHelpers.maybe_put_created_by(role_params, conn)
    case Auth.create_role(params) do
      {:ok, role} ->
        conn
        |> put_flash(:info, "Role created successfully.")
        |> redirect(to: ~p"/roles/#{role}")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :new, changeset: changeset)
    end
  end

  def show(conn, %{"id" => id}) do
    role = Auth.get_role!(id)
    render(conn, :show, role: role)
  end

  def edit(conn, %{"id" => id}) do
    role = Auth.get_role!(id)
    changeset = Auth.change_role(role)
    render(conn, :edit, role: role, changeset: changeset)
  end

  def update(conn, %{"id" => id, "role" => role_params}) do
    role = Auth.get_role!(id)
    params = AuditHelpers.maybe_put_updated_by(role_params, conn)
    case Auth.update_role(role, params) do
      {:ok, role} ->
        conn
        |> put_flash(:info, "Role updated successfully.")
        |> redirect(to: ~p"/roles/#{role}")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :edit, role: role, changeset: changeset)
    end
  end

  def delete(conn, %{"id" => id}) do
    role = Auth.get_role!(id)
    {:ok, _role} = Auth.delete_role(role)

    conn
    |> put_flash(:info, "Role deleted successfully.")
    |> redirect(to: ~p"/roles")
  end

  def permissions(conn, %{"id" => role_id}) do
    role = Auth.get_role!(role_id)
    permissions = Auth.get_permissions_with_assignment_status_for_role(role_id)
    {assigned, unassigned} = Enum.split_with(permissions, & &1.assigned)
    render(conn, "permissions.html", role: role, assigned_permissions: assigned, unassigned_permissions: unassigned)
  end

  def assign_permission(conn, %{"id" => role_id, "permission_id" => permission_id}) do
    case Auth.assign_permission_to_role(role_id, permission_id) do
      {:ok, _} ->
        conn
        |> put_flash(:info, "Permission assigned successfully.")
        |> redirect(to: ~p"/roles/#{role_id}/permissions")
      {:error, :already_assigned} ->
        conn
        |> put_flash(:error, "Permission was already assigned.")
        |> redirect(to: ~p"/roles/#{role_id}/permissions")
      {:error, _} ->
        conn
        |> put_flash(:error, "Could not assign permission.")
        |> redirect(to: ~p"/roles/#{role_id}/permissions")
    end
  end

  def revoke_permission(conn, %{"id" => role_id, "permission_id" => permission_id}) do
    case Auth.revoke_permission_from_role(role_id, permission_id) do
      {:ok, _} ->
        conn
        |> put_flash(:info, "Permission revoked successfully.")
        |> redirect(to: ~p"/roles/#{role_id}/permissions")
      {:error, _} ->
        conn
        |> put_flash(:error, "Could not revoke permission.")
        |> redirect(to: ~p"/roles/#{role_id}/permissions")
    end
  end
end
