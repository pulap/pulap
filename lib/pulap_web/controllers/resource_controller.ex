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
end
