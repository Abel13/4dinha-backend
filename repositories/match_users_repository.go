package repositories

import (
	"4dinha-backend/models"
	"github.com/supabase-community/supabase-go"
)

type MatchUsersRepository struct{}

func NewMatchUsersRepository() *MatchUsersRepository {
	return &MatchUsersRepository{}
}

func (r *MatchUsersRepository) IsDealer(client *supabase.Client, matchID string, playerID string) (bool, error) {
	var matchUser models.MatchUsers

	_, err := client.
		From("match_users").
		Select("*", "", false).
		Eq("match_id", matchID).
		Eq("user_id", playerID).
		Single().
		ExecuteTo(&matchUser)

	if err != nil {
		return false, err
	}

	return matchUser.Dealer, nil
}

func (r *MatchUsersRepository) GetAlivePlayers(client *supabase.Client, matchID string) ([]models.MatchUsers, error) {
	var alivePlayers []models.MatchUsers

	_, err := client.From("match_users").Select("*", "", false).Eq("match_id", matchID).ExecuteTo(&alivePlayers)

	return alivePlayers, err
}
