package auth

import "time"

type Permission string

const (
	PermSystemAdmin Permission = "system:admin"

	PermUsersRead   Permission = "users:read"
	PermUsersWrite  Permission = "users:write"
	PermUsersDelete Permission = "users:delete"

	PermRolesRead   Permission = "roles:read"
	PermRolesWrite  Permission = "roles:write"
	PermRolesDelete Permission = "roles:delete"
	PermRolesManage Permission = "roles:manage"

	PermGrantsRead   Permission = "grants:read"
	PermGrantsWrite  Permission = "grants:write"
	PermGrantsDelete Permission = "grants:delete"
	PermGrantsManage Permission = "grants:manage"

	PermEstatesRead   Permission = "estates:read"
	PermEstatesWrite  Permission = "estates:write"
	PermEstatesDelete Permission = "estates:delete"
	PermEstatesManage Permission = "estates:manage"

	PermEstateItemsRead  Permission = "estates:items:read"
	PermEstateItemsWrite Permission = "estates:items:write"

	PermEstateTagsRead  Permission = "estates:tags:read"
	PermEstateTagsWrite Permission = "estates:tags:write"
)

type PermissionCategory struct {
	Name        string
	Permissions []PermissionInfo
}

type PermissionInfo struct {
	Code        Permission
	Name        string
	Description string
}

var PermissionRegistry = []PermissionCategory{
	{
		Name: "System",
		Permissions: []PermissionInfo{
			{PermSystemAdmin, "System Administrator", "Full system access"},
		},
	},
	{
		Name: "Users",
		Permissions: []PermissionInfo{
			{PermUsersRead, "Read Users", "View user information"},
			{PermUsersWrite, "Write Users", "Create and update users"},
			{PermUsersDelete, "Delete Users", "Remove users from the system"},
		},
	},
	{
		Name: "Roles",
		Permissions: []PermissionInfo{
			{PermRolesRead, "Read Roles", "View role definitions"},
			{PermRolesWrite, "Write Roles", "Create and update roles"},
			{PermRolesDelete, "Delete Roles", "Remove roles"},
			{PermRolesManage, "Manage Roles", "Full role management"},
		},
	},
	{
		Name: "Access Control",
		Permissions: []PermissionInfo{
			{PermGrantsRead, "Read Grants", "View permission assignments"},
			{PermGrantsWrite, "Write Grants", "Assign permissions to users"},
			{PermGrantsDelete, "Delete Grants", "Revoke permissions"},
			{PermGrantsManage, "Manage Grants", "Full grant management"},
		},
	},
	{
		Name: "Estates",
		Permissions: []PermissionInfo{
			{PermEstatesRead, "Read Estates", "View estate information"},
			{PermEstatesWrite, "Write Estates", "Create and update estates"},
			{PermEstatesDelete, "Delete Estates", "Remove estates"},
			{PermEstatesManage, "Manage Estates", "Full estate management"},
			{PermEstateItemsRead, "Read Estate Items", "View estate items"},
			{PermEstateItemsWrite, "Write Estate Items", "Manage estate items"},
			{PermEstateTagsRead, "Read Estate Tags", "View estate tags"},
			{PermEstateTagsWrite, "Write Estate Tags", "Manage estate tags"},
		},
	},
}

func AllPermissions() []Permission {
	var perms []Permission
	for _, cat := range PermissionRegistry {
		for _, p := range cat.Permissions {
			perms = append(perms, p.Code)
		}
	}
	return perms
}

func AllPermissionStrings() []string {
	perms := AllPermissions()
	result := make([]string, len(perms))
	for i, p := range perms {
		result[i] = string(p)
	}
	return result
}

func GetPermissionInfo(code Permission) *PermissionInfo {
	for _, cat := range PermissionRegistry {
		for _, p := range cat.Permissions {
			if p.Code == code {
				return &p
			}
		}
	}
	return nil
}

func EvaluatePermissions(grants []Grant, roles []Role, permission string, scope Scope, now time.Time) bool {
	for _, grant := range grants {
		if grant.ExpiresAt != nil && grant.ExpiresAt.Before(now) {
			continue
		}

		if !ScopeMatches(grant.Scope, scope) {
			continue
		}

		if grant.GrantType == GrantTypePermission && grant.Value == permission {
			return true
		}

		if grant.GrantType == GrantTypeRole {
			rolePermissions := GetRolePermissions(roles, grant.Value)
			if ContainsPermission(rolePermissions, permission) {
				return true
			}
		}
	}

	return false
}

func ScopeMatches(grantScope, requestScope Scope) bool {
	if grantScope.Type == "global" {
		return true
	}

	if grantScope.Type != requestScope.Type {
		return false
	}

	return grantScope.ID == requestScope.ID
}

func GetRolePermissions(roles []Role, roleID string) []string {
	for _, role := range roles {
		if role.ID.String() == roleID {
			return role.Permissions
		}
	}
	return nil
}

func ContainsPermission(permissions []string, permission string) bool {
	for _, p := range permissions {
		if p == permission {
			return true
		}
	}
	return false
}

func FilterValidGrants(grants []Grant, now time.Time) []Grant {
	var validGrants []Grant
	for _, grant := range grants {
		if grant.ExpiresAt == nil || grant.ExpiresAt.After(now) {
			validGrants = append(validGrants, grant)
		}
	}
	return validGrants
}

func GetUserPermissions(grants []Grant, roles []Role, scope Scope, now time.Time) []string {
	var permissions []string
	seen := make(map[string]bool)

	validGrants := FilterValidGrants(grants, now)

	for _, grant := range validGrants {
		if !ScopeMatches(grant.Scope, scope) {
			continue
		}

		if grant.GrantType == GrantTypePermission {
			if !seen[grant.Value] {
				permissions = append(permissions, grant.Value)
				seen[grant.Value] = true
			}
		}

		if grant.GrantType == GrantTypeRole {
			rolePermissions := GetRolePermissions(roles, grant.Value)
			for _, perm := range rolePermissions {
				if !seen[perm] {
					permissions = append(permissions, perm)
					seen[perm] = true
				}
			}
		}
	}

	return permissions
}

func GetEffectiveScopes(grants []Grant, now time.Time) []Scope {
	var scopes []Scope
	seen := make(map[string]bool)

	validGrants := FilterValidGrants(grants, now)

	for _, grant := range validGrants {
		scopeKey := grant.Scope.Type + ":" + grant.Scope.ID
		if !seen[scopeKey] {
			scopes = append(scopes, grant.Scope)
			seen[scopeKey] = true
		}
	}

	return scopes
}

func IsGlobalAdmin(grants []Grant, roles []Role, now time.Time) bool {
	globalScope := Scope{Type: "global", ID: ""}
	return EvaluatePermissions(grants, roles, "system:admin", globalScope, now)
}

func CanManageRole(grants []Grant, roles []Role, targetRoleID string, scope Scope, now time.Time) bool {
	return EvaluatePermissions(grants, roles, "roles:manage", scope, now)
}

func CanGrantPermission(grants []Grant, roles []Role, targetPermission string, scope Scope, now time.Time) bool {
	manageGrantsPermission := EvaluatePermissions(grants, roles, "grants:manage", scope, now)
	if !manageGrantsPermission {
		return false
	}

	hasTargetPermission := EvaluatePermissions(grants, roles, targetPermission, scope, now)
	return hasTargetPermission
}
