defmodule Pulap.Auth do
  @moduledoc """
  The Auth context.
  """

  import Ecto.Query, warn: false
  alias Pulap.Repo

  alias Pulap.Auth.Role

  @doc """
  Returns the list of roles.

  ## Examples

      iex> list_roles()
      [%Role{}, ...]

  """
  def list_roles do
    Repo.all(Role)
  end

  @doc """
  Gets a single role.

  Raises `Ecto.NoResultsError` if the Role does not exist.

  ## Examples

      iex> get_role!(123)
      %Role{}

      iex> get_role!(456)
      ** (Ecto.NoResultsError)

  """
  def get_role!(id), do: Repo.get!(Role, id)

  @doc """
  Creates a role.

  ## Examples

      iex> create_role(%{field: value})
      {:ok, %Role{}}

      iex> create_role(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_role(attrs \\ %{}) do
    %Role{}
    |> Role.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a role.

  ## Examples

      iex> update_role(role, %{field: new_value})
      {:ok, %Role{}}

      iex> update_role(role, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_role(%Role{} = role, attrs) do
    role
    |> Role.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a role.

  ## Examples

      iex> delete_role(role)
      {:ok, %Role{}}

      iex> delete_role(role)
      {:error, %Ecto.Changeset{}}

  """
  def delete_role(%Role{} = role) do
    Repo.delete(role)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking role changes.

  ## Examples

      iex> change_role(role)
      %Ecto.Changeset{data: %Role{}}

  """
  def change_role(%Role{} = role, attrs \\ %{}) do
    Role.changeset(role, attrs)
  end

  def get_roles_with_assignment_status_for_user(user_id) do
    import Ecto.Query
    query =
      from r in Pulap.Auth.Role,
        left_join: ur in "users_roles", on: ur.role_id == r.id and ur.user_id == ^user_id,
        select: %{
          id: r.id,
          name: r.name,
          description: r.description,
          assigned: not is_nil(ur.user_id)
        }
    Pulap.Repo.all(query)
  end

  def assign_role_to_user(user_id, role_id) do
    sql = "INSERT OR IGNORE INTO users_roles (user_id, role_id) VALUES (?, ?)"
    result = Ecto.Adapters.SQL.query!(Pulap.Repo, sql, [user_id, role_id])
    case result.num_rows do
      0 -> {:error, :already_assigned}
      n -> {:ok, n}
    end
  end

  def revoke_role_from_user(user_id, role_id) do
    sql = "DELETE FROM users_roles WHERE user_id = ? AND role_id = ?"
    result = Ecto.Adapters.SQL.query!(Pulap.Repo, sql, [user_id, role_id])
    case result.num_rows do
      0 -> {:error, :not_found}
      n -> {:ok, n}
    end
  end

  alias Pulap.Auth.Permission

  @doc """
  Returns the list of permissions.

  ## Examples

      iex> list_permissions()
      [%Permission{}, ...]

  """
  def list_permissions do
    Repo.all(Permission)
  end

  @doc """
  Gets a single permission.

  Raises `Ecto.NoResultsError` if the Permission does not exist.

  ## Examples

      iex> get_permission!(123)
      %Permission{}

      iex> get_permission!(456)
      ** (Ecto.NoResultsError)

  """
  def get_permission!(id), do: Repo.get!(Permission, id)

  @doc """
  Creates a permission.

  ## Examples

      iex> create_permission(%{field: value})
      {:ok, %Permission{}}

      iex> create_permission(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_permission(attrs \\ %{}) do
    %Permission{}
    |> Permission.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a permission.

  ## Examples

      iex> update_permission(permission, %{field: new_value})
      {:ok, %Permission{}}

      iex> update_permission(permission, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_permission(%Permission{} = permission, attrs) do
    permission
    |> Permission.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a permission.

  ## Examples

      iex> delete_permission(permission)
      {:ok, %Permission{}}

      iex> delete_permission(permission)
      {:error, %Ecto.Changeset{}}

  """
  def delete_permission(%Permission{} = permission) do
    Repo.delete(permission)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking permission changes.

  ## Examples

      iex> change_permission(permission)
      %Ecto.Changeset{data: %Permission{}}

  """
  def change_permission(%Permission{} = permission, attrs \\ %{}) do
    Permission.changeset(permission, attrs)
  end

  def get_permissions_with_assignment_status_for_user(user_id) do
    import Ecto.Query
    query =
      from p in Pulap.Auth.Permission,
        left_join: up in "user_permissions", on: up.permission_id == p.id and up.user_id == ^user_id,
        select: %{
          id: p.id,
          name: p.name,
          description: p.description,
          assigned: not is_nil(up.user_id)
        }
    Pulap.Repo.all(query)
  end

  def assign_permission_to_user(user_id, permission_id) do
    sql = "INSERT OR IGNORE INTO user_permissions (user_id, permission_id) VALUES (?, ?)"
    result = Ecto.Adapters.SQL.query!(Pulap.Repo, sql, [user_id, permission_id])
    case result.num_rows do
      1 -> {:ok, :assigned}
      0 -> {:error, :already_assigned}
    end
  end

  def revoke_permission_from_user(user_id, permission_id) do
    sql = "DELETE FROM user_permissions WHERE user_id = ? AND permission_id = ?"
    result = Ecto.Adapters.SQL.query!(Pulap.Repo, sql, [user_id, permission_id])
    case result.num_rows do
      1 -> {:ok, :revoked}
      0 -> {:error, :not_assigned}
    end
  end

  def get_permissions_with_assignment_status_for_role(role_id) do
    import Ecto.Query
    query =
      from p in Pulap.Auth.Permission,
        left_join: rp in "role_permissions", on: rp.permission_id == p.id and rp.role_id == ^role_id,
        select: %{
          id: p.id,
          name: p.name,
          description: p.description,
          assigned: not is_nil(rp.role_id)
        }
    Pulap.Repo.all(query)
  end

  def assign_permission_to_role(role_id, permission_id) do
    sql = "INSERT OR IGNORE INTO role_permissions (role_id, permission_id) VALUES (?, ?)"
    result = Ecto.Adapters.SQL.query!(Pulap.Repo, sql, [role_id, permission_id])
    case result.num_rows do
      1 -> {:ok, :assigned}
      0 -> {:error, :already_assigned}
    end
  end

  def revoke_permission_from_role(role_id, permission_id) do
    sql = "DELETE FROM role_permissions WHERE role_id = ? AND permission_id = ?"
    result = Ecto.Adapters.SQL.query!(Pulap.Repo, sql, [role_id, permission_id])
    case result.num_rows do
      1 -> {:ok, :revoked}
      0 -> {:error, :not_assigned}
    end
  end

  def get_permissions_with_assignment_status_for_resource(resource_id) do
    import Ecto.Query
    query =
      from p in Pulap.Auth.Permission,
        left_join: rp in "resource_permissions", on: rp.permission_id == p.id and rp.resource_id == ^resource_id,
        select: %{
          id: p.id,
          name: p.name,
          description: p.description,
          assigned: not is_nil(rp.resource_id)
        }
    Pulap.Repo.all(query)
  end

  def assign_permission_to_resource(resource_id, permission_id) do
    sql = "INSERT OR IGNORE INTO resource_permissions (resource_id, permission_id) VALUES (?, ?)"
    result = Ecto.Adapters.SQL.query!(Pulap.Repo, sql, [resource_id, permission_id])
    case result.num_rows do
      1 -> {:ok, :assigned}
      0 -> {:error, :already_assigned}
    end
  end

  def revoke_permission_from_resource(resource_id, permission_id) do
    sql = "DELETE FROM resource_permissions WHERE resource_id = ? AND permission_id = ?"
    result = Ecto.Adapters.SQL.query!(Pulap.Repo, sql, [resource_id, permission_id])
    case result.num_rows do
      1 -> {:ok, :revoked}
      0 -> {:error, :not_assigned}
    end
  end

  alias Pulap.Auth.Resource

  @doc """
  Returns the list of resources.

  ## Examples

      iex> list_resources()
      [%Resource{}, ...]

  """
  def list_resources do
    Repo.all(Resource)
  end

  @doc """
  Gets a single resource.

  Raises `Ecto.NoResultsError` if the Resource does not exist.

  ## Examples

      iex> get_resource!(123)
      %Resource{}

      iex> get_resource!(456)
      ** (Ecto.NoResultsError)

  """
  def get_resource!(id), do: Repo.get!(Resource, id)

  @doc """
  Creates a resource.

  ## Examples

      iex> create_resource(%{field: value})
      {:ok, %Resource{}}

      iex> create_resource(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_resource(attrs \\ %{}) do
    %Resource{}
    |> Resource.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a resource.

  ## Examples

      iex> update_resource(resource, %{field: new_value})
      {:ok, %Resource{}}

      iex> update_resource(resource, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_resource(%Resource{} = resource, attrs) do
    resource
    |> Resource.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a resource.

  ## Examples

      iex> delete_resource(resource)
      {:ok, %Resource{}}

      iex> delete_resource(resource)
      {:error, %Ecto.Changeset{}}

  """
  def delete_resource(%Resource{} = resource) do
    Repo.delete(resource)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking resource changes.

  ## Examples

      iex> change_resource(resource)
      %Ecto.Changeset{data: %Resource{}}

  """
  def change_resource(%Resource{} = resource, attrs \\ %{}) do
    Resource.changeset(resource, attrs)
  end
end
