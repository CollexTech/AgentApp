package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateRoleRequest struct {
	RoleName    string `json:"role_name" binding:"required"`
	Description string `json:"description"`
}

func CreateRole(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	var roleInput CreateRoleRequest
	if err := c.ShouldBindJSON(&roleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := services.CreateRole(env, roleInput.RoleName, roleInput.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

type UpdateRoleRequest struct {
	RoleName    string `json:"role_name" binding:"required"`
	Description string `json:"description"`
}

func UpdateRole(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	roleIDStr := c.Param("id")
	roleID, err := strconv.ParseUint(roleIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	var roleInput UpdateRoleRequest
	if err := c.ShouldBindJSON(&roleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateRole(env, roleID, roleInput.RoleName, roleInput.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role updated successfully"})
}

func DeleteRole(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	roleIDStr := c.Param("id")
	roleID, err := strconv.ParseUint(roleIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	if err := services.DeleteRole(env, roleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}

type AssignRoleToUserRequest struct {
	UserID   string   `json:"user_id" binding:"required"`
	RoleList []string `json:"role_list" binding:"required"`
}

func AssignRoleToUser(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	var assignInput AssignRoleToUserRequest

	if err := c.ShouldBindJSON(&assignInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AssignRolesToUser(env, assignInput.UserID, assignInput.RoleList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Roles assigned successfully"})
}

type RemoveRoleFromUserRequest struct {
	UserID string `json:"user_id" binding:"required"`
	RoleID string `json:"role_id" binding:"required"`
}

func RemoveRoleFromUser(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	var removeInput RemoveRoleFromUserRequest

	if err := c.ShouldBindJSON(&removeInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.RemoveRoleFromUser(env, removeInput.UserID, removeInput.RoleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role removed successfully"})
}

func GetRolesByUser(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	userIDStr := c.Param("user_id")

	roles, err := services.GetRolesByUser(env, userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func ListAllRoles(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	roles, err := services.ListAllRoles(env)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func GetMyRoles(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	roles, err := services.GetRolesByUser(env, env.AuthDtos.User.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func GetMyPermissions(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	c.JSON(http.StatusOK, env.PermissionList)
}
