package services

import (
	"backend/models"
	"backend/repository"
)

func ListCases(env *models.Env, status string, agencyID string) ([]models.Case, error) {
	agencyRepository := repository.NewAgencyRepository(env.DbConn)

	cases, err := agencyRepository.ListCases(status, agencyID)
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
