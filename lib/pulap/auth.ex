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
        where: r.contextual == false,
        left_join: ur in "users_roles",
        on: ur.role_id == r.id and ur.user_id == ^user_id and is_nil(ur.context_type),
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
        left_join: up in "user_permissions",
        on: up.permission_id == p.id and up.user_id == ^user_id,
        left_join: ur in "users_roles",
        on: ur.user_id == ^user_id,
        left_join: r in Pulap.Auth.Role,
        on: r.id == ur.role_id,
        left_join: rp in "role_permissions",
        on: rp.role_id == ur.role_id and rp.permission_id == p.id,
        group_by: [p.id, p.name, p.description, up.user_id],
        select: %{
          id: p.id,
          name: p.name,
          description: p.description,
          direct: not is_nil(up.user_id),
          indirect: fragment("count(DISTINCT ?) > 0", rp.role_id),
          source_roles:
            fragment("group_concat(DISTINCT ?) FILTER (WHERE ? IS NOT NULL)", r.name, rp.role_id)
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
        left_join: rp in "role_permissions",
        on: rp.permission_id == p.id and rp.role_id == ^role_id,
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
        left_join: rp in "resource_permissions",
        on: rp.permission_id == p.id and rp.resource_id == ^resource_id,
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

  alias Pulap.Org.Organization

  @doc """
  Returns the list of organizations.
  """
  def list_organizations do
    Repo.all(Organization)
  end

  @doc """
  Gets a single organization.
  Raises `Ecto.NoResultsError` if the Organization does not exist.
  """
  def get_organization!(id), do: Repo.get!(Organization, id)

  @doc """
  Gets a single organization.
  """
  def get_organization(id), do: Repo.get(Organization, id)

  @doc """
  Creates an organization.
  """
  def create_organization(attrs \\ %{}) do
    %Organization{}
    |> Organization.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates an organization.
  """
  def update_organization(%Organization{} = organization, attrs) do
    organization
    |> Organization.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes an organization.
  """
  def delete_organization(%Organization{} = organization) do
    Repo.delete(organization)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking organization changes.
  """
  def change_organization(%Organization{} = organization, attrs \\ %{}) do
    Organization.changeset(organization, attrs)
  end

  alias Pulap.Org.Team

  @doc """
  Returns the list of teams.

  ## Examples

      iex> list_teams()
      [%Team{}, ...]

  """
  def list_teams do
    Repo.all(Team)
  end

  @doc """
  Gets a single team.

  Raises `Ecto.NoResultsError` if the Team does not exist.

  ## Examples

      iex> get_team!(123)
      %Team{}

      iex> get_team!(456)
      ** (Ecto.NoResultsError)

  """
  def get_team!(id), do: Repo.get!(Team, id)

  @doc """
  Creates a team.

  ## Examples

      iex> create_team(%{field: value})
      {:ok, %Team{}}

      iex> create_team(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_team(attrs \\ %{}) do
    %Team{}
    |> Team.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a team.

  ## Examples

      iex> update_team(team, %{field: new_value})
      {:ok, %Team{}}

      iex> update_team(team, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_team(%Team{} = team, attrs) do
    team
    |> Team.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a team.

  ## Examples

      iex> delete_team(team)
      {:ok, %Team{}}

      iex> delete_team(team)
      {:error, %Ecto.Changeset{}}

  """
  def delete_team(%Team{} = team) do
    Repo.delete(team)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking team changes.

  ## Examples

      iex> change_team(team)
      %Ecto.Changeset{data: %Team{}}

  """
  def change_team(%Team{} = team, attrs \\ %{}) do
    Team.changeset(team, attrs)
  end
end
