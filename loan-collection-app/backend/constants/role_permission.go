package constants

var RolePermissionsMap = map[string][]string{
	"admin": {
		"view_users",
		"create_user",
		"update_user",
		"delete_user",
		"view_roles",
		"create_role",
		"update_role",
		"delete_role",
		"assign_role_to_user",
		"remove_role_from_user",
		"view_users",
	},
	"agent": {
		"view_cases",
		"view_trails",
		"generate_payment_link",
	},
	"default": {
		"view_my_permissions",
	},
}
