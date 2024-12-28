package models

import (
	"time"

	"gorm.io/datatypes"
)

type User struct {
	ID           string          `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
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

type Role struct {
	ID          string    `gorm:"column:id;primaryKey;autoIncrement"`
	RoleName    string    `gorm:"column:role_name;unique;not null"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type UserRoleMap struct {
	ID         string    `gorm:"column:id;primaryKey;autoIncrement"`
	UserID     string    `gorm:"column:user_id;not null"`
	RoleID     string    `gorm:"column:role_id;not null"`
	IsActive   bool      `gorm:"column:is_active;default:true"`
	AssignedAt time.Time `gorm:"column:assigned_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
	User       User      `gorm:"foreignKey:UserID"`
	Role       Role      `gorm:"foreignKey:RoleID"`
}

func (Role) TableName() string {
	return "roles"
}

func (UserRoleMap) TableName() string {
	return "user_role_map"
}
