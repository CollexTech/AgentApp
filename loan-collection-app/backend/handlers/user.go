package handlers

import (
	"backend/constants"
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

func CreateUserHandler(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err := env.Validator.Struct(req); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := services.CreateUser(env, req.Username, req.Password)
	if err != nil {
		env.Logger.Error(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user_id": user.ID,
	})
}

type ListUserResponse struct {
	Id       string   `json:"id"`
	Username string   `json:"username"`
	Email    *string  `json:"email"`
	IsActive bool     `json:"is_active"`
	Role     []string `json:"role_list"`
}

func ListAllUsers(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	users, err := services.ListAllUsers(env)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	usersResponse := make([]ListUserResponse, len(users))
	for i, user := range users {
		roles, err := services.GetRolesByUser(env, user.ID)
		if err != nil {
			env.Logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		roleList := make([]string, len(roles))
		for j, role := range roles {
			roleList[j] = role.RoleName
		}
		usersResponse[i] = ListUserResponse{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			IsActive: user.IsActive,
			Role:     roleList,
		}
	}

	c.JSON(http.StatusOK, usersResponse)
}
