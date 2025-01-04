package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AgencyCaseResponse struct {
	Id                     string    `json:"id"`
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

func GetAgencyCasesHandler(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	cases, err := services.ListCases(env, env.AuthDtos.User.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := []AgencyCaseResponse{}
	for _, c := range cases {
		response = append(response, AgencyCaseResponse{
			Id:                     c.ID,
			LoanID:                 c.LoanID,
			ExternalCustomerID:     c.ExternalCustomerID,
			EMIAmount:              c.EMIAmount,
			PrincipalOutstanding:   c.PrincipalOutstanding,
			InterestOutstanding:    c.InterestOutstanding,
			CaseStatus:             c.CaseStatus,
			EMIDate:                c.EMIDate,
			DPDBucket:              c.DPDBucket,
			DPD:                    c.DPD,
			DisbursalDate:          c.DisbursalDate,
			InsuranceActive:        c.InsuranceActive,
			LoanDescription:        c.LoanDescription,
			EMIsPaidTillDate:       c.EMIsPaidTillDate,
			EMIsPending:            c.EMIsPending,
			BounceCharges:          c.BounceCharges,
			NachPresentationStatus: c.NachPresentationStatus,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

type AssignCaseRequest struct {
	CaseID string `json:"case_id" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}

func AssignAgencyCaseHandler(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	var req AssignCaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mapping := &models.CaseUserMap{
		CaseID: req.CaseID,
		UserID: req.UserID,
	}

	if err := services.AssignCaseToUser(env, mapping); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Case assigned successfully"})
}

func GetAgencyUsersHandler(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	users, err := services.ListAgencyUsers(env, env.AuthDtos.User.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}
