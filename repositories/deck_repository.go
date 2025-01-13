package repositories

import (
	"4dinha-backend/models"
	"fmt"
	"github.com/supabase-community/supabase-go"
)

type DeckRepository struct{}

func NewDeckRepository() *DeckRepository {
	return &DeckRepository{}
}

func (r *DeckRepository) GetAllCards(client *supabase.Client) ([]models.Deck, error) {
	var cards []models.Deck

	_, err := client.From("deck").Select("*", "", false).ExecuteTo(&cards)

	if err != nil {
		return nil, fmt.Errorf("error to get cards: %w", err)
	}

	return cards, nil
}
