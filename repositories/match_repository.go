package repositories

import (
	"4dinha-backend/models"
	"github.com/supabase-community/supabase-go"
)

type MatchRepository struct{}

func NewMatchRepository() *MatchRepository {
	return &MatchRepository{}
}

func (r *MatchRepository) GetMatch(client *supabase.Client, matchID string) (models.Matches, error) {
	var match models.Matches

	_, err := client.
		From("matches").
		Select("*", "", false).
		Eq("id", matchID).
		Single().
		ExecuteTo(&match)

	return match, err
}
