package utils

import "backend/models"

func HasPermission(env *models.Env, permission string) bool {
	permissionsList := env.PermissionList
	for _, val := range permissionsList {
		if val == permission {
			return true
		}
	}
	return false
}
