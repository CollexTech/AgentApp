package repository

import (
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

type AgencyRepository struct {
	db *gorm.DB
}

func NewAgencyRepository(db *gorm.DB) *AgencyRepository {
	return &AgencyRepository{db: db}
}

func (r *AgencyRepository) ListAllAgencies() ([]models.Agency, error) {
	var agencies []models.Agency
	result := r.db.Where("status = ?", "ACTIVE").Find(&agencies)
	if result.Error != nil {
		return nil, result.Error
	}
	return agencies, nil
}

func (r *AgencyRepository) CreateAgency(agency *models.Agency) error {
	if agency.Status == "" {
		agency.Status = "ACTIVE"
	}
	result := r.db.Create(agency)
	return result.Error
}

func (r *AgencyRepository) DeleteAgency(agencyID string) error {
	result := r.db.Delete(&models.Agency{}, "id = ?", agencyID)
	return result.Error
}

func (r *AgencyRepository) AssignUserToAgency(mapping *models.AgencyUserMap) error {
	// Check if user and agency exist
	var user models.User
	var agency models.Agency

	if err := r.db.First(&user, "id = ?", mapping.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	if err := r.db.First(&agency, "id = ?", mapping.AgencyID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("agency not found")
		}
		return err
	}

	result := r.db.Create(mapping)
	return result.Error
}

func (r *AgencyRepository) AssignCaseToUser(mapping *models.CaseUserMap) error {
	// Check if user and case exist
	var user models.User
	if err := r.db.First(&user, "id = ?", mapping.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Assuming you have a Case model
	var case_ models.Case
	if err := r.db.First(&case_, "id = ?", mapping.CaseID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("case not found")
		}
		return err
	}

	result := r.db.Create(mapping)
	return result.Error
}

func (r *AgencyRepository) CreateAgencyCaseMap(agencyCaseMapList []models.AgencyCaseMap) error {
	result := r.db.Create(agencyCaseMapList)
	return result.Error
}

func (r *AgencyRepository) ListCases(status string, agencyID string) ([]models.Case, error) {
	cases := []models.Case{}
	result := r.db.Where("status = ?", status).Where("agency_id = ? and status = ?", agencyID, status).Find(&cases)
	return cases, result.Error
}
