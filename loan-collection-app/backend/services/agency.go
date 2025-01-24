package services

import (
	"backend/models"
	"backend/repository"
)

func ListAllAgencies(env *models.Env) ([]models.Agency, error) {
	repo := repository.NewAgencyRepository(env.DbConn)
	return repo.ListAllAgencies()
}

func CreateAgency(env *models.Env, agency *models.Agency) error {
	if agency.Status == "" {
		agency.Status = "ACTIVE"
	}

	repo := repository.NewAgencyRepository(env.DbConn)
	return repo.CreateAgency(agency)
}

func DeleteAgency(env *models.Env, agencyID string) error {
	repo := repository.NewAgencyRepository(env.DbConn)
	return repo.DeleteAgency(agencyID)
}

func AssignUserToAgency(env *models.Env, mapping *models.AgencyUserMap) error {
	repo := repository.NewAgencyRepository(env.DbConn)
	return repo.AssignUserToAgency(mapping)
}

func AssignCaseToUser(env *models.Env, mapping *models.CaseUserMap) error {
	repo := repository.NewAgencyRepository(env.DbConn)
	return repo.AssignCaseToUser(mapping)
}
func ListAgencyUsers(env *models.Env, agencyID string) ([]models.AgencyUserDetails, error) {
	repo := repository.NewAgencyRepository(env.DbConn)
	return repo.ListAgencyUsers(agencyID)
}

func ListMyAgencyUsers(env *models.Env, userID string) ([]models.AgencyUserDetails, error) {
	userRepository := repository.NewUserRepository(env.DbConn)
	agencyID, err := userRepository.GetUserAgencyID(userID)
	if err != nil {
		return nil, err
	}
	agencyRepository := repository.NewAgencyRepository(env.DbConn)
	return agencyRepository.ListAgencyUsers(agencyID)
}

func ListUnassignedUsers(env *models.Env) ([]models.User, error) {
	repo := repository.NewAgencyRepository(env.DbConn)
	return repo.ListUnassignedUsers()
}
