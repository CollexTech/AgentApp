package services

import (
	"backend/models"
	"backend/repository"
	"errors"
	"strconv"
	"time"
)

func CreateCasesFromCSV(env *models.Env, records [][]string) ([]models.Case, error) {
	if len(records) < 2 { // Check if there's at least a header and one data row
		return nil, errors.New("CSV file is empty or missing data rows")
	}

	// Skip header row (first row)
	var cases []models.Case
	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) < 16 { // We expect 16 columns as per the CSV structure
			continue
		}

		// Parse numeric values
		emiAmount, _ := strconv.ParseFloat(record[2], 64)
		principalOutstanding, _ := strconv.ParseFloat(record[3], 64)
		interestOutstanding, _ := strconv.ParseFloat(record[4], 64)
		dpd, _ := strconv.Atoi(record[8])
		emisPaidTillDate, _ := strconv.Atoi(record[12])
		emisPending, _ := strconv.Atoi(record[13])
		bounceCharges, _ := strconv.ParseFloat(record[14], 64)

		// Parse dates
		emiDate, _ := time.Parse("2006-01-02", record[6])
		disbursalDate, _ := time.Parse("2006-01-02", record[9])

		// Parse boolean
		insuranceActive := record[10] == "true"

		// Create Case struct
		case_ := models.Case{
			LoanID:                 record[0],
			ExternalCustomerID:     record[1],
			EMIAmount:              emiAmount,
			PrincipalOutstanding:   principalOutstanding,
			InterestOutstanding:    interestOutstanding,
			CaseStatus:             record[5],
			EMIDate:                emiDate,
			DPDBucket:              record[7],
			DPD:                    dpd,
			DisbursalDate:          disbursalDate,
			InsuranceActive:        insuranceActive,
			LoanDescription:        record[11],
			EMIsPaidTillDate:       emisPaidTillDate,
			EMIsPending:            emisPending,
			BounceCharges:          bounceCharges,
			NachPresentationStatus: record[15],
			CreatedAt:              time.Now(),
			UpdatedAt:              time.Now(),
		}
		cases = append(cases, case_)
	}

	repo := repository.NewCaseRepository(env.DbConn)
	return repo.CreateCases(cases)
}

func GetUnassignedCases(env *models.Env) ([]models.Case, error) {
	repo := repository.NewCaseRepository(env.DbConn)
	return repo.GetUnassignedCases()
}

func AssignCasesToAgency(env *models.Env, agencyID string, caseIDs []string) error {
	repo := repository.NewCaseRepository(env.DbConn)
	return repo.AssignCasesToAgency(agencyID, caseIDs)
}

func GetAssignedUserByCaseID(env *models.Env, caseID string) (*models.User, error) {
	repo := repository.NewCaseRepository(env.DbConn)
	return repo.GetAssignedUserByCaseID(caseID)
}

func GetAssignedCases(env *models.Env, userID string) ([]models.Case, error) {
	repo := repository.NewCaseRepository(env.DbConn)
	return repo.GetAssignedCases(userID)
}

func GetCaseDetails(env *models.Env, caseID string) (*models.Case, *models.User, error) {
	repo := repository.NewCaseRepository(env.DbConn)
	caseData, err := repo.GetCase(caseID)
	if err != nil {
		return nil, nil, err
	}
	userData, err := repo.GetAssignedUserByCaseID(caseID)
	if err != nil {
		return nil, nil, err
	}
	return caseData, userData, nil
}
