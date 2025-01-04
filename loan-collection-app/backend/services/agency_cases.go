package services

import (
	"backend/models"
	"backend/repository"
)

func ListCases(env *models.Env, userID string) ([]models.Case, error) {
	agencyRepository := repository.NewAgencyRepository(env.DbConn)
	userRepository := repository.NewUserRepository(env.DbConn)
	agencyID, err := userRepository.GetUserAgencyID(userID)
	if err != nil {
		return nil, err
	}

	cases, err := agencyRepository.ListCases(agencyID)
	if err != nil {
		return nil, err
	}
	return cases, nil
}

func AssignCases(env *models.Env, agencyID string, caseIDs []string) ([]models.AgencyCaseMap, error) {
	agencyRepository := repository.NewAgencyRepository(env.DbConn)

	agencyCaseMapList := []models.AgencyCaseMap{}
	for _, caseID := range caseIDs {
		agencyCaseMapList = append(agencyCaseMapList, models.AgencyCaseMap{
			AgencyID: agencyID,
			CaseID:   caseID,
		})
	}
	err := agencyRepository.CreateAgencyCaseMap(agencyCaseMapList)
	if err != nil {
		return nil, err
	}

	return agencyCaseMapList, nil
}
