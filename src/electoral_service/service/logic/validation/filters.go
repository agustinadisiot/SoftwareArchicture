package validation

import (
	"electoral_service/models"
	"fmt"
	p_f "pipes_and_filters"
	"time"
)

func GetAvailableFilters() map[string]p_f.FilterWithParams {

	availableFilters := map[string]p_f.FilterWithParams{
		"validate_election_date":              FilterValidateDate,
		"validate_party_list":                 FilterValidatePoliticalPartyList,
		"validate_voter_candidate_list":       FilterValidateCandidateList,
		"validate_unique_party_per_candidate": FilterValidateUniquePartyPerCandidate,
		"validate_election_mode":              FilterElectionMode,
	}
	return availableFilters
}

func FilterValidatePoliticalPartyList(data any, params map[string]any) error {
	election := data.(models.ElectionModelEssential)
	if len(election.PoliticalParties) == 0 {
		return fmt.Errorf("politicalPartyList is nil")
	}
	return nil
}

func FilterValidateCandidateList(data any, params map[string]any) error {
	election := data.(models.ElectionModelEssential)
	for _, party := range election.PoliticalParties {
		if len(party.Candidates) == 0 {
			return fmt.Errorf("candidateList is nil")
		}
	}
	if len(election.Voters) == 0 {
		return fmt.Errorf("voterList is nil")
	}
	return nil
}

func FilterValidateUniquePartyPerCandidate(data any, params map[string]any) error {
	election := data.(models.ElectionModelEssential)
	candidatesToCheck := make(map[string]bool)
	for _, party := range election.PoliticalParties {
		for _, candidate := range party.Candidates {
			if _, ok := candidatesToCheck[candidate.Id]; ok {
				return fmt.Errorf("candidate %s is in more than one party", candidate.Id)
			} else {
				candidatesToCheck[candidate.Id] = true
			}
		}
	}
	return nil
}

func FilterValidateDate(data any, params map[string]any) error {
	election := data.(models.ElectionModelEssential)
	StartingDate, _ := time.Parse(time.RFC3339, election.StartingDate)
	EndDate, _ := time.Parse(time.RFC3339, election.FinishingDate)

	if StartingDate.After(EndDate) || StartingDate.Equal(EndDate) {
		return fmt.Errorf("starting date is after end date")
	}
	return nil
}

func FilterElectionMode(data any, params map[string]any) error {
	election := data.(models.ElectionModelEssential)
	if election.ElectionMode != "unico" && election.ElectionMode != "multi" {
		return fmt.Errorf("election mode is not valid")
	}
	return nil
}