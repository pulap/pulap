package auth

type Permission string

const (
	// System Administration
	PermSystemAdmin  Permission = "system:admin"
	PermSystemConfig Permission = "system:config"

	// Authentication Management
	PermUsersRead   Permission = "users:read"
	PermUsersWrite  Permission = "users:write"
	PermUsersDelete Permission = "users:delete"

	// Authorization Management
	PermRolesRead   Permission = "roles:read"
	PermRolesWrite  Permission = "roles:write"
	PermRolesDelete Permission = "roles:delete"
	PermRolesManage Permission = "roles:manage"

	PermGrantsRead   Permission = "grants:read"
	PermGrantsWrite  Permission = "grants:write"
	PermGrantsDelete Permission = "grants:delete"
	PermGrantsManage Permission = "grants:manage"
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
			{PermSystemConfig, "System Configuration", "Manage system settings"},
		},
	},
	{
		Name: "Authentication",
		Permissions: []PermissionInfo{
			{PermUsersRead, "Read Users", "View user information"},
			{PermUsersWrite, "Write Users", "Create and update users"},
			{PermUsersDelete, "Delete Users", "Remove users from the system"},
		},
	},
	{
		Name: "Authorization",
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
