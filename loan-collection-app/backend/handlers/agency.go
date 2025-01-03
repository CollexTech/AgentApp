package handlers

import (
	"backend/models"
	"backend/services"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListAgenciesResponse struct {
	Id         string `json:"id"`
	AgencyName string `json:"agency_name"`
	Status     string `json:"status"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

func ListAllAgencies(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	agencies, err := services.ListAllAgencies(env)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := []ListAgenciesResponse{}
	for _, agency := range agencies {
		var agencyDetails map[string]string
		json.Unmarshal(agency.AgencyDetails, &agencyDetails)
		response = append(response, ListAgenciesResponse{
			Id:         agency.ID,
			AgencyName: agency.AgencyName,
			Status:     agency.Status,
			Address:    agencyDetails["address"],
			Phone:      agencyDetails["phone"],
			Email:      agencyDetails["email"],
		})
	}

	c.JSON(http.StatusOK, response)
}

type CreateAgencyRequest struct {
	AgencyName string `json:"agency_name" binding:"required"`
	Status     string `json:"status" binding:"required"`
	Address    string `json:"address" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Email      string `json:"email" binding:"required"`
}

func CreateAgency(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	var agencyRequest CreateAgencyRequest
	if err := c.ShouldBindJSON(&agencyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	agencyDetails := map[string]string{
		"address": agencyRequest.Address,
		"phone":   agencyRequest.Phone,
		"email":   agencyRequest.Email,
	}
	agencyDetailsJSON, _ := json.Marshal(agencyDetails)

	agency := models.Agency{
		AgencyName:    agencyRequest.AgencyName,
		Status:        agencyRequest.Status,
		AgencyDetails: agencyDetailsJSON,
	}

	if err := services.CreateAgency(env, &agency); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, agency)
}

func DeleteAgency(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	agencyID := c.Param("id")

	if err := services.DeleteAgency(env, agencyID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

type AssignUserToAgencyRequest struct {
	UserID     string  `json:"user_id" binding:"required"`
	AgencyID   string  `json:"agency_id" binding:"required"`
	AgencyRole string  `json:"agency_role" binding:"required"`
	ManagerID  *string `json:"manager_id"`
}

func AssignUserToAgency(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	var mappingRequest AssignUserToAgencyRequest
	if err := c.ShouldBindJSON(&mappingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mapping := models.AgencyUserMap{
		UserID:     mappingRequest.UserID,
		AgencyID:   mappingRequest.AgencyID,
		AgencyRole: mappingRequest.AgencyRole,
		ManagerID:  mappingRequest.ManagerID,
		IsActive:   true,
	}

	if err := services.AssignUserToAgency(env, &mapping); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapping)
}

func AssignCaseToUser(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	var mapping models.CaseUserMap
	if err := c.ShouldBindJSON(&mapping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AssignCaseToUser(env, &mapping); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapping)
}

type ListAgencyUsersResponse struct {
	ID        string  `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Role      string  `json:"role"`
	ManagerID *string `json:"manager_id,omitempty"`
}

type UnassignedUsersResponse struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Email    *string `json:"email"`
}

func GetAgencyUsers(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	agencyID := c.Param("agency_id")
	users, err := services.ListAgencyUsers(env, agencyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := []ListAgencyUsersResponse{}
	for _, user := range users {
		response = append(response, ListAgencyUsersResponse{
			ID:        user.UserID,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.AgencyRole,
			ManagerID: user.ManagerID,
		})
	}

	c.JSON(http.StatusOK, response)
}

func GetUnassignedUsers(c *gin.Context) {
	val, _ := c.Get("env")
	env := val.(*models.Env)

	users, err := services.ListUnassignedUsers(env)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := []UnassignedUsersResponse{}
	for _, user := range users {
		response = append(response, UnassignedUsersResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	c.JSON(http.StatusOK, response)
}
