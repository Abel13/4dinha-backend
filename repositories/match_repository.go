package repositories

import (
	"4dinha-backend/models"
	"fmt"
	"github.com/supabase-community/supabase-go"
)

type MatchRepository struct{}

func NewMatchRepository() *MatchRepository {
	return &MatchRepository{}
}

func (r *MatchRepository) GetMatch(client *supabase.Client, matchID string) ([]models.Matches, error) {
	var matchList []models.Matches

	matches, _, err := client.From("matches").Select("*", "", false).Eq("id", matchID).Execute()

	fmt.Println(matches)
	return matchList, err
}
