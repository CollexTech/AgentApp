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
		"view_agencies",
		"create_agency",
		"assign_agency_user",
		"assign_case",
		"delete_agency",
		"view_agency_users",
		"view_unassigned_users",
		"assign_user_to_agency",
		"view_agency_user_mapping",
		"assign_user_to_agency",
		"upload_cases",
		"view_unassigned_cases",
		"assign_cases",
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
