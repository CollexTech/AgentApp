package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Each trail is a record of communication attempts
type Trail struct {
	TrailID     int    `json:"trail_id"`
	CaseID      int    `json:"case_id"`
	Contacted   bool   `json:"contacted"`
	PaymentDate string `json:"payment_date"` // If user says they'll pay on a date
	Remarks     string `json:"remarks"`
}

var trails []Trail
var trailIDCounter = 1

// POST /api/cases/:caseID/trails
func AddTrail(c *gin.Context) {
	caseIDStr := c.Param("caseID")
	caseID, _ := strconv.Atoi(caseIDStr)

	var newTrail Trail
	if err := c.ShouldBindJSON(&newTrail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trail data"})
		return
	}

	newTrail.TrailID = trailIDCounter
	newTrail.CaseID = caseID
	trailIDCounter++

	trails = append(trails, newTrail)
	c.JSON(http.StatusOK, newTrail)
}

// GET /api/cases/:caseID/trails
func GetTrails(c *gin.Context) {
	caseIDStr := c.Param("caseID")
	caseID, _ := strconv.Atoi(caseIDStr)

	var caseTrails []Trail
	for _, t := range trails {
		if t.CaseID == caseID {
			caseTrails = append(caseTrails, t)
		}
	}

	c.JSON(http.StatusOK, caseTrails)
}

// GET /api/cases/:caseID/payment-link
func GetPaymentLink(c *gin.Context) {
	caseIDStr := c.Param("caseID")
	// Just return a dummy link
	paymentLink := "https://payment.example.com/pay?caseID=" + caseIDStr
	c.JSON(http.StatusOK, gin.H{"payment_link": paymentLink})
}
