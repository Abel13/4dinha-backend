package repositories

import (
	"4dinha-backend/models"
	"fmt"
	"github.com/supabase-community/supabase-go"
)

type MatchUsersRepository interface {
	IsDealer(matchID, playerID string) (bool, error)
	GetAlivePlayers(matchID string) ([]models.MatchUsers, error)
}

type SupabaseMatchUsersRepository struct {
	DB *supabase.Client
}

func (r *SupabaseMatchUsersRepository) IsDealer(matchID string, playerID string) (bool, error) {
	var matchUser []models.MatchUsers
	
	_, err := r.DB.
		From("match_users").
		Select("*", "", false).
		ExecuteTo(&matchUser)

	if err != nil {
		return false, err
	}

	if matchUser[0].ID != "" {
		return true, nil
	}
	return false, nil
}

func (r *SupabaseMatchUsersRepository) GetAlivePlayers(matchID string) ([]models.MatchUsers, error) {
	var alivePlayers []models.MatchUsers

	matchUsers, _, err := r.DB.From("match_users").Select("*", "", false).Eq("match_id", matchID).ExecuteString()
	fmt.Println(matchUsers)

	return alivePlayers, err
}
