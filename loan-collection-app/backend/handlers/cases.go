package handlers

import (
	"net/http"
	"strconv"

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
func GetAgentCases(c *gin.Context) {
	// You can calculate total earnings for the day, or pass back a separate struct
	totalEarningsForDay := 500.0

	c.JSON(http.StatusOK, gin.H{
		"cases":          cases,
		"total_earnings": totalEarningsForDay,
	})
}

// GET /api/cases/:caseID
func GetCaseDetails(c *gin.Context) {
	caseIDStr := c.Param("caseID")
	caseID, _ := strconv.Atoi(caseIDStr)

	for _, kase := range cases {
		if kase.CaseID == caseID {
			c.JSON(http.StatusOK, kase)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Case not found"})
}
