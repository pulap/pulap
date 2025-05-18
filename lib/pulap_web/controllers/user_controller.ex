defmodule PulapWeb.UserController do
  use PulapWeb, :controller

  plug :put_layout, html: {PulapWeb.Layouts, :auth}

  alias Pulap.Accounts
  alias Pulap.Accounts.User
  alias PulapWeb.AuditHelpers

  def index(conn, _params) do
    users = Accounts.list_users()
    render(conn, :index, users: users)
  end

  def new(conn, _params) do
    changeset = Accounts.change_user(%User{})
    render(conn, :new, changeset: changeset)
  end

  def create(conn, %{"user" => user_params}) do
    params = AuditHelpers.maybe_put_created_by(user_params, conn)
    case Accounts.create_user(params) do
      {:ok, user} ->
        conn
        |> put_flash(:info, "User created successfully.")
        |> redirect(to: ~p"/users/#{user}")

      {:error, %Ecto.Changeset{} = changeset} ->
        IO.inspect(changeset.errors, label: "[DEBUG] User creation failed")
        render(conn, :new, changeset: changeset)
    end
  end

  def show(conn, %{"id" => id}) do
    user = Accounts.get_user!(id)
    render(conn, :show, user: user)
  end

  def edit(conn, %{"id" => id}) do
    user = Accounts.get_user!(id)
    changeset = Accounts.change_user(user)
    render(conn, :edit, user: user, changeset: changeset)
  end

  def update(conn, %{"id" => id, "user" => user_params}) do
    user = Accounts.get_user!(id)
    params = AuditHelpers.maybe_put_updated_by(user_params, conn)

    case Accounts.update_user(user, params) do
      {:ok, user} ->
        conn
        |> put_flash(:info, "User updated successfully.")
        |> redirect(to: ~p"/users/#{user}")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :edit, user: user, changeset: changeset)
    end
  end

  def delete(conn, %{"id" => id}) do
    user = Accounts.get_user!(id)
    {:ok, _user} = Accounts.delete_user(user)

    conn
    |> put_flash(:info, "User deleted successfully.")
    |> redirect(to: ~p"/users")
  end

  def roles(conn, %{"id" => user_id}) do
    user = Pulap.Accounts.get_user!(user_id)
    roles = Pulap.Auth.get_roles_with_assignment_status_for_user(user_id)
    {assigned, unassigned} = Enum.split_with(roles, & &1.assigned)
    render(conn, "roles.html", user: user, assigned_roles: assigned, unassigned_roles: unassigned)
  end

  def assign_role(conn, %{"id" => user_id, "role_id" => role_id}) do
    case Pulap.Auth.assign_role_to_user(user_id, role_id) do
      {:ok, _} ->
        conn
        |> put_flash(:info, "Role assigned successfully.")
        |> redirect(to: ~p"/users/#{user_id}/roles")
      {:error, :already_assigned} ->
        conn
        |> put_flash(:error, "Role was already assigned.")
        |> redirect(to: ~p"/users/#{user_id}/roles")
      {:error, _} ->
        conn
        |> put_flash(:error, "Could not assign role.")
        |> redirect(to: ~p"/users/#{user_id}/roles")
    end
  end

  def revoke_role(conn, %{"id" => user_id, "role_id" => role_id}) do
    case Pulap.Auth.revoke_role_from_user(user_id, role_id) do
      {:ok, _} ->
        conn
        |> put_flash(:info, "Role revoked successfully.")
        |> redirect(to: ~p"/users/#{user_id}/roles")
      {:error, _} ->
        conn
        |> put_flash(:error, "Could not revoke role.")
        |> redirect(to: ~p"/users/#{user_id}/roles")
    end
  end

  def permissions(conn, %{"id" => user_id}) do
    user = Pulap.Accounts.get_user!(user_id)
    permissions = Pulap.Auth.get_permissions_with_assignment_status_for_user(user_id)
    assigned = Enum.filter(permissions, fn p ->
      p.direct == true or (p.indirect == 1 and p.source_roles not in [nil, ""])
    end)
    unassigned = Enum.reject(permissions, fn p ->
      p.direct == true or (p.indirect == 1 and p.source_roles not in [nil, ""])
    end)
    render(conn, "permissions.html", user: user, assigned_permissions: assigned, unassigned_permissions: unassigned)
  end

  def assign_permission(conn, %{"id" => user_id, "permission_id" => permission_id}) do
    case Pulap.Auth.assign_permission_to_user(user_id, permission_id) do
      {:ok, _} ->
        conn
        |> put_flash(:info, "Permission assigned successfully.")
        |> redirect(to: ~p"/users/#{user_id}/permissions")
      {:error, :already_assigned} ->
        conn
        |> put_flash(:error, "Permission was already assigned.")
        |> redirect(to: ~p"/users/#{user_id}/permissions")
      {:error, _} ->
        conn
        |> put_flash(:error, "Could not assign permission.")
        |> redirect(to: ~p"/users/#{user_id}/permissions")
    end
  end

  def revoke_permission(conn, %{"id" => user_id, "permission_id" => permission_id}) do
    case Pulap.Auth.revoke_permission_from_user(user_id, permission_id) do
      {:ok, _} ->
        conn
        |> put_flash(:info, "Permission revoked successfully.")
        |> redirect(to: ~p"/users/#{user_id}/permissions")
      {:error, _} ->
        conn
        |> put_flash(:error, "Could not revoke permission.")
        |> redirect(to: ~p"/users/#{user_id}/permissions")
    end
  end
end
