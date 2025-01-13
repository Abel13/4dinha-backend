package repositories

import (
	"4dinha-backend/models"
	"github.com/supabase-community/supabase-go"
)

type PlayerCardsRepository struct {
}

func NewPlayerCardsRepository() *PlayerCardsRepository {
	return &PlayerCardsRepository{}
}

func (r *PlayerCardsRepository) CreatePlayerCards(client *supabase.Client, playerCards []models.PlayerCards) error {
	_, _, err := client.
		From("player_cards").
		Insert(playerCards, false, "", "minimal", "").
		Execute()

	if err != nil {
		return err
	}

	return nil
}
