package repositories

import (
	"4dinha-backend/models"
	"github.com/supabase-community/supabase-go"
)

type DeckRepository struct{}

func NewDeckRepository() *DeckRepository {
	return &DeckRepository{}
}

func (r *DeckRepository) GetAllCards(client *supabase.Client) []models.Deck {
	var cards []models.Deck

	_, err := client.From("deck").Select("*", "", false).ExecuteTo(&cards)

	if err != nil {
		return nil
	}

	return cards
}

func (r *DeckRepository) GetCard(client *supabase.Client, cardID string) models.Deck {
	var card models.Deck

	_, err := client.
		From("deck").
		Select("*", "", false).Eq("id", cardID).Single().ExecuteTo(&card)

	if err != nil {
		return card
	}

	return card
}
