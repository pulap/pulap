defmodule PulapWeb.RoleController do
  use PulapWeb, :controller

  alias Pulap.Auth
  alias Pulap.Auth.Role

  def index(conn, _params) do
    roles = Auth.list_roles()
    render(conn, :index, roles: roles)
  end

  def new(conn, _params) do
    changeset = Auth.change_role(%Role{})
    render(conn, :new, changeset: changeset)
  end

  def create(conn, %{"role" => role_params}) do
    case Auth.create_role(role_params) do
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

    case Auth.update_role(role, role_params) do
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
end
