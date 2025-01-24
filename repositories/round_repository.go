package repositories

import (
	"4dinha-backend/models"
	"fmt"
	"github.com/supabase-community/supabase-go"
)

type RoundRepository struct {
}

func NewRoundRepo() *RoundRepository {
	return &RoundRepository{}
}

func (r *RoundRepository) CreateRound(client *supabase.Client, matchID string, roundNumber int, shuffledCards *[]models.Deck) error {
	cardID := (*shuffledCards)[0].ID
	*shuffledCards = (*shuffledCards)[1:]

	var round = models.Rounds{
		MatchID:     matchID,
		RoundNumber: roundNumber,
		Status:      models.StatusBetting,
		Trump:       cardID,
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

func (r *RoundRepository) CurrentRound(client *supabase.Client, roundNumber string, matchID string) (*models.Rounds, error) {
	var round *models.Rounds

	_, err := client.
		From("rounds").
		Select("*", "", false).
		Eq("match_id", matchID).
		Eq("round_number", roundNumber).
		Single().
		ExecuteTo(&round)

	if err != nil {
		return nil, fmt.Errorf("error to get cards: %w", err)
	}

	return round, nil
}

func (r *RoundRepository) GetTrumpsByPower(client *supabase.Client, trumpPower string) ([]models.Deck, error) {
	var trumps []models.Deck

	_, err := client.
		From("deck").
		Select("*", "", false).
		Eq("power", trumpPower).
		ExecuteTo(&trumps)

	return trumps, err
}
