package handlers

import (
	"backend/models"
	"backend/services"
	"encoding/csv"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadCasesHandler handles the CSV upload and case creation
func UploadCasesHandler(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
		return
	}
	defer openedFile.Close()

	reader := csv.NewReader(openedFile)
	records, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing CSV"})
		return
	}

	cases, err := services.CreateCasesFromCSV(env, records)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cases uploaded successfully",
		"count":   len(cases),
	})
}

type GetUnassignedCasesResponse struct {
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

// GetUnassignedCasesHandler returns cases not assigned to any agency
func GetUnassignedCasesHandler(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	cases, err := services.GetUnassignedCases(env)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := []GetUnassignedCasesResponse{}
	for _, c := range cases {
		response = append(response, GetUnassignedCasesResponse{
			ID:                     c.ID,
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

// AssignCasesRequest represents the request body for case assignment
type AssignCasesRequest struct {
	AgencyID string   `json:"agency_id" binding:"required"`
	CaseIDs  []string `json:"case_ids" binding:"required"`
}

// AssignCasesHandler handles assigning cases to an agency
func AssignCasesHandler(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	var req AssignCasesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.AssignCasesToAgency(env, req.AgencyID, req.CaseIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cases assigned successfully"})
}
