package models

import (
	"time"
)

type Case struct {
	ID                     string    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	LoanID                 string    `gorm:"type:varchar(50)"`
	ExternalCustomerID     string    `gorm:"type:varchar(50)"`
	EMIAmount              float64   `gorm:"type:numeric(10,2)"`
	PrincipalOutstanding   float64   `gorm:"type:numeric(10,2)"`
	InterestOutstanding    float64   `gorm:"type:numeric(10,2)"`
	CaseStatus             string    `gorm:"type:varchar(50)"`
	EMIDate                time.Time `gorm:"type:date"`
	DPDBucket              string    `gorm:"type:varchar(50)"`
	DPD                    int       `gorm:"type:integer"`
	DisbursalDate          time.Time `gorm:"type:date"`
	InsuranceActive        bool      `gorm:"type:boolean;default:false"`
	LoanDescription        string    `gorm:"type:text"`
	EMIsPaidTillDate       int       `gorm:"type:integer"`
	EMIsPending            int       `gorm:"type:integer"`
	BounceCharges          float64   `gorm:"type:numeric(10,2)"`
	NachPresentationStatus string    `gorm:"type:varchar(50)"`
	CreatedAt              time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt              time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (Case) TableName() string {
	return "case"
}
