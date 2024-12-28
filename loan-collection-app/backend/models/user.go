package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type User struct {
	ID           uuid.UUID       `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username     string          `gorm:"type:varchar(50);not null;unique"`
	Email        *string         `gorm:"type:varchar(100)"`
	PasswordHash string          `gorm:"type:text;not null"`
	ProfileData  *datatypes.JSON `gorm:"type:jsonb"`
	IsActive     bool            `gorm:"default:true"`
	CreatedAt    time.Time       `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time       `gorm:"default:CURRENT_TIMESTAMP"`
}

// TableName sets the table name for the User model
func (User) TableName() string {
	return "users"
}
