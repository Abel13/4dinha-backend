package repositories

import (
	"4dinha-backend/models"

	"github.com/supabase-community/postgrest-go"
	"github.com/supabase-community/supabase-go"
)

type MatchActionRepository struct {
}

func NewMatchActionRepository() *MatchActionRepository {
	return &MatchActionRepository{}
}

func (r *MatchActionRepository) RegisterAction(client *supabase.Client, matchID string, roundNumber int, action models.Actions) error {
	var newAction = models.MatchActionsInput{
		MatchID:     matchID,
		Action:      action,
		RoundNumber: roundNumber,
	}

	_, _, err := client.
		From("match_actions").
		Insert(newAction, false, "", "minimal", "").
		Execute()

	if err != nil {
		return err
	}

	return nil
}

func (r *MatchActionRepository) GetLastAction(client *supabase.Client, matchID string) (models.MatchActions, error) {
	var action models.MatchActions

	_, err := client.
		From("match_actions").
		Select("*", "", false).
		Eq("match_id", matchID).
		Order("created_at", &postgrest.OrderOpts{
			Ascending: false,
		}).
		Limit(1, "").
		Single().
		ExecuteTo(&action)

	return action, err
}
