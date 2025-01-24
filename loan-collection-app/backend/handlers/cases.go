package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// In-memory data
type Case struct {
	CaseID        int     `json:"case_id"`
	UserName      string  `json:"user_name"`
	LoanID        string  `json:"loan_id"`
	LoanAmount    float64 `json:"loan_amount"`
	EMIMonthly    float64 `json:"emi_monthly"`
	DaysPastDue   int     `json:"days_past_due"`
	CustomerAddr  string  `json:"customer_addr"`
	CustomerPhone string  `json:"customer_phone"`
	// ... add other fields as needed
}

var cases = []Case{
	{
		CaseID: 1, UserName: "John Doe", LoanID: "LN123456", LoanAmount: 15000, EMIMonthly: 1500, DaysPastDue: 10,
		CustomerAddr: "123 Main St, Cityville", CustomerPhone: "9876543210",
	},
	{
		CaseID: 2, UserName: "Jane Smith", LoanID: "LN789012", LoanAmount: 20000, EMIMonthly: 2000, DaysPastDue: 5,
		CustomerAddr: "456 Elm St, Townsville", CustomerPhone: "9898989898",
	},
}

// GET /api/cases
type GetAgentCasesResponse struct {
	ID                     string    `json:"id"`
	LoanID                 string    `json:"loan_id"`
	ExternalCustomerID     string    `json:"external_customer_id"`
	EMIAmount              float64   `json:"emi_amount"`
	PrincipalOutstanding   float64   `json:"principal_outstanding"`
	InterestOutstanding    float64   `json:"interest_outstanding"`
	CaseStatus             string    `json:"case_status"`
	EMIDate                time.Time `json:"emi_date"`
	DPDBucket              string    `json:"dpd_bucket"`
	DPD                    int       `json:"dpd"`
	DisbursalDate          time.Time `json:"disbursal_date"`
	InsuranceActive        bool      `json:"insurance_active"`
	LoanDescription        string    `json:"loan_description"`
	EMIsPaidTillDate       int       `json:"emis_paid_till_date"`
	EMIsPending            int       `json:"emis_pending"`
	BounceCharges          float64   `json:"bounce_charges"`
	NachPresentationStatus string    `json:"nach_presentation_status"`
}

func GetAgentCases(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	cases, err := services.GetAssignedCases(env, env.AuthDtos.User.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := []GetAgentCasesResponse{}
	for _, caseData := range cases {
		response = append(response, GetAgentCasesResponse{
			ID:                     caseData.ID,
			LoanID:                 caseData.LoanID,
			ExternalCustomerID:     caseData.ExternalCustomerID,
			EMIAmount:              caseData.EMIAmount,
			PrincipalOutstanding:   caseData.PrincipalOutstanding,
			InterestOutstanding:    caseData.InterestOutstanding,
			CaseStatus:             caseData.CaseStatus,
			EMIDate:                caseData.EMIDate,
			DPDBucket:              caseData.DPDBucket,
			DPD:                    caseData.DPD,
			DisbursalDate:          caseData.DisbursalDate,
			InsuranceActive:        caseData.InsuranceActive,
			LoanDescription:        caseData.LoanDescription,
			EMIsPaidTillDate:       caseData.EMIsPaidTillDate,
			EMIsPending:            caseData.EMIsPending,
			BounceCharges:          caseData.BounceCharges,
			NachPresentationStatus: caseData.NachPresentationStatus,
		})
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

type GetCaseDetailsResponse struct {
	CaseID                 string    `json:"case_id"`
	AgentName              string    `json:"agent_name"`
	LoanID                 string    `json:"loan_id"`
	LoanAmount             float64   `json:"loan_amount"`
	EMIMonthly             float64   `json:"emi_monthly"`
	DaysPastDue            int       `json:"days_past_due"`
	CustomerAddr           string    `json:"customer_addr"`
	CustomerPhone          string    `json:"customer_phone"`
	CaseStatus             string    `json:"case_status"`
	EMIDate                time.Time `json:"emi_date"`
	DPDBucket              string    `json:"dpd_bucket"`
	DPD                    int       `json:"dpd"`
	DisbursalDate          time.Time `json:"disbursal_date"`
	InsuranceActive        bool      `json:"insurance_active"`
	LoanDescription        string    `json:"loan_description"`
	EMIsPaidTillDate       int       `json:"emis_paid_till_date"`
	EMIsPending            int       `json:"emis_pending"`
	BounceCharges          float64   `json:"bounce_charges"`
	NachPresentationStatus string    `json:"nach_presentation_status"`
}

// GET /api/cases/:caseID
func GetCaseDetails(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	caseIDStr := c.Param("caseID")

	caseData, userData, err := services.GetCaseDetails(env, caseIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := GetCaseDetailsResponse{
		CaseID:                 caseData.ID,
		AgentName:              userData.Username,
		LoanID:                 caseData.LoanID,
		LoanAmount:             caseData.PrincipalOutstanding,
		EMIMonthly:             caseData.EMIAmount,
		DaysPastDue:            caseData.DPD,
		CustomerAddr:           "123 Main St, Cityville",
		CustomerPhone:          "9876543210",
		CaseStatus:             caseData.CaseStatus,
		EMIDate:                caseData.EMIDate,
		DPDBucket:              caseData.DPDBucket,
		DPD:                    caseData.DPD,
		DisbursalDate:          caseData.DisbursalDate,
		InsuranceActive:        caseData.InsuranceActive,
		LoanDescription:        caseData.LoanDescription,
		EMIsPaidTillDate:       caseData.EMIsPaidTillDate,
		EMIsPending:            caseData.EMIsPending,
		BounceCharges:          caseData.BounceCharges,
		NachPresentationStatus: caseData.NachPresentationStatus,
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}
