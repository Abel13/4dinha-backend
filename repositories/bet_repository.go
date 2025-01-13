package repositories

import (
	"4dinha-backend/models"
	"github.com/supabase-community/supabase-go"
)

type BetRepository struct {
}

func NewBetRepository() *BetRepository {
	return &BetRepository{}
}

func (r *BetRepository) GetRoundBets(client *supabase.Client, matchID string, roundNumber string) ([]models.Bets, error) {
	var bets []models.Bets

	_, err := client.
		From("bets").
		Select("*", "", false).
		Eq("match_id", matchID).
		Eq("round_number", roundNumber).
		ExecuteTo(&bets)

	return bets, err
}
