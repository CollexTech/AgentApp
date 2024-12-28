package middlewares

import (
	"backend/constants"
	"backend/models"
	"backend/repository"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// AuthMiddleware validates JWT token and attaches user information to the context
func AuthMiddleware(c *gin.Context) {
	// Get environment from context
	val, exists := c.Get("env")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Environment not configured"})
		return
	}
	env := val.(*models.Env)

	// Extract authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	// Check Bearer token format
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	// Extract token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse and validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Get JWT secret from configuration
		jwtSecret := viper.GetString(constants.JWT_SECRET)
		if jwtSecret == "" {
			return nil, fmt.Errorf("JWT secret is not configured")
		}

		return []byte(jwtSecret), nil
	})

	// Check for token parsing errors
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Validate token claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract user ID from claims
		userID, ok := claims["user_id"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
			return
		}

		// Find user in database to ensure they exist and are active
		userRepo := repository.NewUserRepository(env.DbConn)
		user, err := userRepo.FindByID(userID)
		if err != nil || user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		env.AuthDtos.User = user

		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
}
