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

	result := r.db.Table("agency_user_map").Create(mapping)
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

func (r *AgencyRepository) ListCases(agencyID string) ([]models.Case, error) {
	cases := []models.Case{}
	result := r.db.Table("cases").
		Joins("JOIN agency_case_map ON cases.id = agency_case_map.case_id").
		Where("agency_case_map.agency_id = ?", agencyID).
		Find(&cases)
	return cases, result.Error
}

func (r *AgencyRepository) ListAgencyUsers(agencyID string) ([]models.AgencyUserDetails, error) {
	var users []models.AgencyUserDetails

	result := r.db.Table("agency_user_map").
		Select("users.id as user_id, users.username, users.email, agency_user_map.agency_role, agency_user_map.manager_id").
		Joins("JOIN users ON users.id = agency_user_map.user_id").
		Where("agency_user_map.agency_id = ? AND agency_user_map.is_active = true", agencyID).
		Scan(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *AgencyRepository) ListUnassignedUsers() ([]models.User, error) {
	var users []models.User

	// Get users who are not assigned to any agency or whose assignments are not active
	result := r.db.Table("users").
		Select("users.*").
		Joins("LEFT JOIN agency_user_map ON users.id = agency_user_map.user_id AND agency_user_map.is_active = true").
		Where("agency_user_map.id IS NULL").
		Scan(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
