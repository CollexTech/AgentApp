package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

// CreateRole creates a new role
func (r *RoleRepository) CreateRole(roleName, description string) (*models.Role, error) {
	role := &models.Role{
		RoleName:    roleName,
		Description: description,
	}
	result := r.db.Create(role)
	return role, result.Error
}

// UpdateRole updates an existing role
func (r *RoleRepository) UpdateRole(roleID uint64, roleName, description string) error {
	return r.db.Model(&models.Role{}).Where("id = ?", roleID).Updates(map[string]interface{}{
		"role_name":   roleName,
		"description": description,
	}).Error
}

// DeleteRole soft deletes a role
func (r *RoleRepository) DeleteRole(roleID uint64) error {
	return r.db.Delete(&models.Role{}, roleID).Error
}

// AssignRoleToUser assigns a role to a user
func (r *RoleRepository) AssignRoleToUser(userID, roleID string) error {
	userRoleMap := &models.UserRoleMap{
		UserID:   userID,
		RoleID:   roleID,
		IsActive: true,
	}
	return r.db.Create(userRoleMap).Error
}

// RemoveRoleFromUser removes a role from a user
func (r *RoleRepository) RemoveRoleFromUser(userID, roleID string) error {
	return r.db.Where("user_id = ? AND role_id = ?", userID, roleID).Delete(&models.UserRoleMap{}).Error
}

// GetRolesByUser retrieves all roles for a specific user
func (r *RoleRepository) GetRolesByUser(userID string) ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Joins("JOIN user_role_map ON user_role_map.role_id = roles.id").
		Where("user_role_map.user_id = ? AND user_role_map.is_active = true", userID).
		Find(&roles).Error
	return roles, err
}

// ListAllRoles retrieves all roles
func (r *RoleRepository) ListAllRoles() ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Find(&roles).Error
	return roles, err
}
