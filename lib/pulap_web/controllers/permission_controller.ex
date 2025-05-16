defmodule PulapWeb.PermissionController do
  use PulapWeb, :controller

  alias Pulap.Auth
  alias Pulap.Auth.Permission

  def index(conn, _params) do
    permissions = Auth.list_permissions()
    render(conn, :index, permissions: permissions)
  end

  def new(conn, _params) do
    changeset = Auth.change_permission(%Permission{})
    render(conn, :new, changeset: changeset)
  end

  def create(conn, %{"permission" => permission_params}) do
    case Auth.create_permission(permission_params) do
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

    case Auth.update_permission(permission, permission_params) do
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
