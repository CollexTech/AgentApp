package handlers

import (
	"backend/constants"
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func LoginHandler(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)
	var credentials LoginRequest
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials format"})
		return
	}

	if err := env.Validator.Struct(credentials); err != nil {
		env.Logger.Error(err.Error())
		response := constants.RESPONSE_BAD_REQUEST
		response["error"] = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := services.Login(env, credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
		"token":   token,
	})
}
