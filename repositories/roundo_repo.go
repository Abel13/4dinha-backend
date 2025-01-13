package repositories

import (
	"4dinha-backend/models"
	"github.com/supabase-community/supabase-go"
)

type RoundRepo struct {
}

func NewRoundRepo() *RoundRepo {
	return &RoundRepo{}
}

func (r *RoundRepo) CreateRound(client *supabase.Client, matchID string, roundNumber int) error {
	var round = models.Rounds{
		MatchID:     matchID,
		RoundNumber: roundNumber,
		Status:      models.StatusDealing,
	}

	_, _, err := client.
		From("rounds").
		Insert(round,
			false,
			"",
			"minimal",
			"",
		).Execute()

	if err != nil {
		return err
	}

	return nil
}
