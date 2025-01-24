package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AgencyCaseResponse struct {
	Id                     string                  `json:"id"`
	LoanID                 string                  `json:"loan_id"`
	ExternalCustomerID     string                  `json:"external_customer_id"`
	EMIAmount              float64                 `json:"emi_amount"`
	PrincipalOutstanding   float64                 `json:"principal_outstanding"`
	InterestOutstanding    float64                 `json:"interest_outstanding"`
	CaseStatus             string                  `json:"case_status"`
	EMIDate                time.Time               `json:"emi_date"`
	DPDBucket              string                  `json:"dpd_bucket"`
	DPD                    int                     `json:"dpd"`
	DisbursalDate          time.Time               `json:"disbursal_date"`
	InsuranceActive        bool                    `json:"insurance_active"`
	LoanDescription        string                  `json:"loan_description"`
	EMIsPaidTillDate       int                     `json:"emis_paid_till_date"`
	EMIsPending            int                     `json:"emis_pending"`
	BounceCharges          float64                 `json:"bounce_charges"`
	NachPresentationStatus string                  `json:"nach_presentation_status"`
	AssignedTo             *AgencyCaseResponseUser `json:"assigned_to"`
}

type AgencyCaseResponseUser struct {
	Id       string  `json:"id"`
	Username string  `json:"username"`
	Email    *string `json:"email"`
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
	for _, caseData := range cases {
		agencyCaseResponse := AgencyCaseResponse{
			Id:                     caseData.ID,
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
		}
		user, err := services.GetAssignedUserByCaseID(env, caseData.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if user != nil {
			agencyCaseResponse.AssignedTo = &AgencyCaseResponseUser{
				Id:       user.ID,
				Username: user.Username,
				Email:    user.Email,
			}
		}
		response = append(response, agencyCaseResponse)
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

func GetMyAgencyUsersHandler(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	users, err := services.ListMyAgencyUsers(env, env.AuthDtos.User.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}
