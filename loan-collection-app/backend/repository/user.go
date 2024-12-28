package repository

import (
	"backend/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) ValidateCredentials(username, password string) (*models.User, error) {
	user, err := r.FindByUsername(username)
	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (r *UserRepository) CreateUser(username, password string) (*models.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create new user
	user := &models.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
	}

	// Save user to database
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) FindByID(userID string) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, "id = ?", userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
