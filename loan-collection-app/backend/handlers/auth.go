package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Dummy user for demonstration
var fakeAgent = map[string]string{
	"username": "agent1",
	"password": "password123",
}

// LoginHandler authenticates an agent and returns a simple token
func LoginHandler(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials format"})
		return
	}

	if credentials.Username == fakeAgent["username"] && credentials.Password == fakeAgent["password"] {
		// In real app, use JWT
		token := "SAMPLE_TOKEN"
		c.JSON(http.StatusOK, gin.H{
			"message": "Logged in successfully",
			"token":   token,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

// Simple Middleware to check token
func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No or invalid token"})
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token != "SAMPLE_TOKEN" { // In real app, validate JWT
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	c.Next()
}
