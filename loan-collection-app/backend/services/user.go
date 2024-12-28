package services

import (
	"backend/constants"
	"backend/models"
	"backend/repository"

	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func Login(env *models.Env, username, password string) (string, error) {
	repository := repository.NewUserRepository(env.DbConn)
	user, err := repository.ValidateCredentials(username, password)
	if err != nil {
		return "", err
	}

	// Generate JWT token
	token, err := generateJWTToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateJWTToken(user *models.User) (string, error) {
	jwtSecret := viper.GetString(constants.JWT_SECRET)
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT secret is not configured")
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func CreateUser(env *models.Env, username, password string) (*models.User, error) {
	repository := repository.NewUserRepository(env.DbConn)

	// Check if user already exists
	existingUser, _ := repository.FindByUsername(username)
	if existingUser != nil {
		return nil, fmt.Errorf("username already exists")
	}

	// Validate password strength (optional, you can add more complex validation)
	if len(password) < 8 {
		return nil, fmt.Errorf("password must be at least 8 characters long")
	}

	// Create user
	user, err := repository.CreateUser(username, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
