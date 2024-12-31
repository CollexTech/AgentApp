package handlers

import (
	"backend/models"
	"backend/services"
	"encoding/csv"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UploadCasesHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Open the file
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
		return
	}
	defer openedFile.Close()

	// Parse CSV
	reader := csv.NewReader(openedFile)
	records, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing CSV"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var cases []models.Case

	// Skip header row
	for i, record := range records {
		if i == 0 {
			continue
		}
		// Parse CSV data into Case struct
		// Add validation and error handling as needed
		loanCase := models.Case{
			ExternalCustomerID: record[0],
			LoanID:             record[1],
			// ... parse other fields
			CaseStatus: "PENDING",
		}
		cases = append(cases, loanCase)
	}

	// Bulk insert cases
	if result := db.Create(&cases); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating cases"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cases uploaded successfully"})
}

func ListCasesHandler(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	status := c.Query("status")
	agencyID := c.Query("agency_id")

	cases, err := services.ListCases(env, status, agencyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching cases"})
		return
	}

	c.JSON(http.StatusOK, cases)

}

type AssignCasesRequest struct {
	AgencyID string   `json:"agency_id"`
	CaseIDs  []string `json:"case_ids"`
}

func AssignCasesHandler(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	var request AssignCasesRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	agencyCaseMapList, err := services.AssignCases(env, request.AgencyID, request.CaseIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error assigning cases"})
		return
	}

	c.JSON(http.StatusOK, agencyCaseMapList)
}
