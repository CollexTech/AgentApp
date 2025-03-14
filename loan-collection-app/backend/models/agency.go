package models

import (
	"encoding/json"
	"time"
)

type Agency struct {
	ID            string          `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	AgencyName    string          `gorm:"type:varchar(255);not null"`
	Status        string          `gorm:"type:varchar(255);not null"`
	AgencyDetails json.RawMessage `gorm:"type:jsonb;not null"`
	CreatedAt     time.Time       `gorm:"type:timestamp;not null"`
	UpdatedAt     time.Time       `gorm:"type:timestamp;not null"`
}

type AgencyUserMap struct {
	ID         string    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	AgencyID   string    `gorm:"type:uuid;not null"`
	UserID     string    `gorm:"type:uuid;not null"`
	ManagerID  *string   `gorm:"type:uuid;null"`
	AgencyRole string    `gorm:"type:varchar(255);not null"`
	IsActive   bool      `gorm:"type:boolean;not null"`
	AssignedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt  time.Time `gorm:"type:timestamp;not null"`
}

func (AgencyUserMap) TableName() string {
	return "agency_user_map"
}

type CaseUserMap struct {
	ID         string    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CaseID     string    `gorm:"type:uuid;not null"`
	UserID     string    `gorm:"type:uuid;not null"`
	AssignedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt  time.Time `gorm:"type:timestamp;not null"`
}

func (CaseUserMap) TableName() string {
	return "case_user_map"
}

type AgencyCaseMap struct {
	ID         string    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	AgencyID   string    `gorm:"type:uuid;not null"`
	CaseID     string    `gorm:"type:uuid;not null"`
	AssignedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt  time.Time `gorm:"type:timestamp;not null"`
}

func (AgencyCaseMap) TableName() string {
	return "agency_case_map"
}

type AgencyUserDetails struct {
	UserID     string  `json:"user_id"`
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	AgencyRole string  `json:"agency_role"`
	ManagerID  *string `json:"manager_id"`
}
