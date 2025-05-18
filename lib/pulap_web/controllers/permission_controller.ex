defmodule PulapWeb.PermissionController do
  use PulapWeb, :controller

  plug :put_layout, html: {PulapWeb.Layouts, :auth}

  alias Pulap.Auth
  alias Pulap.Auth.Permission
  alias PulapWeb.AuditHelpers

  def index(conn, _params) do
    permissions = Auth.list_permissions()
    render(conn, :index, permissions: permissions)
  end

  def new(conn, _params) do
    changeset = Auth.change_permission(%Permission{})
    render(conn, :new, changeset: changeset)
  end

  def create(conn, %{"permission" => permission_params}) do
    params = AuditHelpers.maybe_put_created_by(permission_params, conn)
    case Auth.create_permission(params) do
      {:ok, permission} ->
        IO.puts("[INFO] Permission created: #{inspect(permission)}")
        conn
        |> put_flash(:info, "Permission created successfully.")
        |> redirect(to: ~p"/permissions/#{permission}")

      {:error, %Ecto.Changeset{} = changeset} ->
        IO.puts("[ERROR] Permission creation failed: #{inspect(changeset.errors)}")
        render(conn, :new, changeset: changeset)
    end
  end

  def show(conn, %{"id" => id}) do
    permission = Auth.get_permission!(id)
    render(conn, :show, permission: permission)
  end

  def edit(conn, %{"id" => id}) do
    permission = Auth.get_permission!(id)
    changeset = Auth.change_permission(permission)
    render(conn, :edit, permission: permission, changeset: changeset)
  end

  def update(conn, %{"id" => id, "permission" => permission_params}) do
    permission = Auth.get_permission!(id)
    params = AuditHelpers.maybe_put_updated_by(permission_params, conn)

    case Auth.update_permission(permission, params) do
      {:ok, permission} ->
        conn
        |> put_flash(:info, "Permission updated successfully.")
        |> redirect(to: ~p"/permissions/#{permission}")

      {:error, %Ecto.Changeset{} = changeset} ->
        render(conn, :edit, permission: permission, changeset: changeset)
    end
  end

  def delete(conn, %{"id" => id}) do
    permission = Auth.get_permission!(id)
    {:ok, _permission} = Auth.delete_permission(permission)

    conn
    |> put_flash(:info, "Permission deleted successfully.")
    |> redirect(to: ~p"/permissions")
  end
end
