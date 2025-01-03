package models

import (
	"time"
)

type Case struct {
	ID                     string    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	LoanID                 string    `gorm:"type:varchar(50);column:loan_id"`
	ExternalCustomerID     string    `gorm:"type:varchar(50);column:external_customer_id"`
	EMIAmount              float64   `gorm:"type:numeric(10,2);column:emi_amount"`
	PrincipalOutstanding   float64   `gorm:"type:numeric(10,2);column:principal_outstanding"`
	InterestOutstanding    float64   `gorm:"type:numeric(10,2);column:interest_outstanding"`
	CaseStatus             string    `gorm:"type:varchar(50);column:case_status"`
	EMIDate                time.Time `gorm:"type:date;column:emi_date"`
	DPDBucket              string    `gorm:"type:varchar(50);column:dpd_bucket"`
	DPD                    int       `gorm:"type:integer;column:dpd"`
	DisbursalDate          time.Time `gorm:"type:date;column:disbursal_date"`
	InsuranceActive        bool      `gorm:"type:boolean;default:false;column:insurance_active"`
	LoanDescription        string    `gorm:"type:text;column:loan_description"`
	EMIsPaidTillDate       int       `gorm:"type:integer;column:emis_paid_till_date"`
	EMIsPending            int       `gorm:"type:integer;column:emis_pending"`
	BounceCharges          float64   `gorm:"type:numeric(10,2);column:bounce_charges"`
	NachPresentationStatus string    `gorm:"type:varchar(50);column:nach_presentation_status"`
	CreatedAt              time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt              time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`
}

func (Case) TableName() string {
	return "cases"
}
