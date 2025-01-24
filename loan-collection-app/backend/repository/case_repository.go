package repository

import (
	"backend/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type CaseRepository struct {
	db *gorm.DB
}

func NewCaseRepository(db *gorm.DB) *CaseRepository {
	return &CaseRepository{db: db}
}

func (r *CaseRepository) CreateCases(cases []models.Case) ([]models.Case, error) {
	result := r.db.Create(&cases)
	if result.Error != nil {
		return nil, result.Error
	}
	return cases, nil
}

func (r *CaseRepository) GetUnassignedCases() ([]models.Case, error) {
	var cases []models.Case

	err := r.db.Table("cases").
		Select("cases.*").
		Joins("LEFT JOIN agency_case_map ON cases.id = agency_case_map.case_id").
		Where("agency_case_map.id IS NULL").
		Find(&cases).Error

	if err != nil {
		return nil, err
	}

	return cases, nil
}

func (r *CaseRepository) AssignCasesToAgency(agencyID string, caseIDs []string) error {
	// Begin transaction
	tx := r.db.Begin()

	// Create agency case mappings
	for _, caseID := range caseIDs {
		mapping := models.AgencyCaseMap{
			AgencyID:   agencyID,
			CaseID:     caseID,
			AssignedAt: time.Now(),
			UpdatedAt:  time.Now(),
		}

		if err := tx.Create(&mapping).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Update case status
	if err := tx.Model(&models.Case{}).
		Where("id IN ?", caseIDs).
		Update("case_status", "ASSIGNED").Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *CaseRepository) GetAssignedUserByCaseID(caseID string) (*models.User, error) {
	var user models.User

	err := r.db.Table("users").
		Select("users.*").
		Joins("JOIN case_user_map ON users.id = case_user_map.user_id").
		Where("case_user_map.case_id = ?", caseID).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *CaseRepository) GetAssignedCases(userID string) ([]models.Case, error) {

	var cases []models.Case

	err := r.db.Table("cases").
		Select("cases.*").
		Joins("JOIN case_user_map ON cases.id = case_user_map.case_id").
		Where("case_user_map.user_id = ?", userID).
		Find(&cases).Error

	if err != nil {
		return nil, err
	}

	return cases, nil
}

func (r *CaseRepository) GetCase(caseID string) (*models.Case, error) {
	var caseData models.Case
	err := r.db.Where("id = ?", caseID).First(&caseData).Error
	if err != nil {
		return nil, err
	}
	return &caseData, nil
}
