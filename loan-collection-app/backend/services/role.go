package services

import (
	"backend/models"
	"backend/repository"
	"errors"
)

// Business logic methods that can include validation, logging, etc.
func CreateRole(env *models.Env, roleName, description string) (*models.Role, error) {
	roleRepo := repository.NewRoleRepository(env.DbConn)
	// Add validation logic
	if roleName == "" {
		return nil, errors.New("role name cannot be empty")
	}
	return roleRepo.CreateRole(roleName, description)
}

func UpdateRole(env *models.Env, roleID uint64, roleName, description string) error {
	if roleName == "" {
		return errors.New("role name cannot be empty")
	}
	roleRepo := repository.NewRoleRepository(env.DbConn)
	return roleRepo.UpdateRole(roleID, roleName, description)
}

func DeleteRole(env *models.Env, roleID uint64) error {
	roleRepo := repository.NewRoleRepository(env.DbConn)
	return roleRepo.DeleteRole(roleID)
}

func AssignRoleToUser(env *models.Env, userID, roleID string) error {
	roleRepo := repository.NewRoleRepository(env.DbConn)
	return roleRepo.AssignRoleToUser(userID, roleID)
}

func RemoveRoleFromUser(env *models.Env, userID, roleID string) error {
	roleRepo := repository.NewRoleRepository(env.DbConn)
	return roleRepo.RemoveRoleFromUser(userID, roleID)
}

func GetRolesByUser(env *models.Env, userID string) ([]models.Role, error) {
	roleRepo := repository.NewRoleRepository(env.DbConn)
	return roleRepo.GetRolesByUser(userID)
}

func ListAllRoles(env *models.Env) ([]models.Role, error) {
	roleRepo := repository.NewRoleRepository(env.DbConn)
	return roleRepo.ListAllRoles()
}
