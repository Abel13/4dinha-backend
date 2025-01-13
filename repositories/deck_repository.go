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

func (r *DeckRepository) GetAllCards(client supabase.Client) ([]models.Deck, error) {
	var cards []models.Deck

	deck, _, err := client.From("deck").Select("*", "", false).Execute()
	fmt.Println(deck)

	if err != nil {
		return nil, fmt.Errorf("erro ao buscar cartas: %w", err)
	}

	return cards, nil
}
