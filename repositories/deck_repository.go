package repositories

import (
	"4dinha-backend/models"
	"fmt"
	"github.com/supabase-community/supabase-go"
)

type DeckRepository interface {
	GetAllCards() ([]models.Deck, error)
}

type SupabaseDeckRepository struct {
	DB *supabase.Client
}

func (r *SupabaseDeckRepository) GetAllCards() ([]models.Deck, error) {
	var cards []models.Deck

	deck, _, err := r.DB.From("deck").Select("*", "", false).Execute()
	fmt.Println(deck)

	if err != nil {
		return nil, fmt.Errorf("erro ao buscar cartas: %w", err)
	}

	return cards, nil
}
