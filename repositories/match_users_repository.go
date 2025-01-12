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
	matchUser, _, err := r.DB.From("match_users").Select("*", "", false).Eq("match_id", matchID).Eq("user_id", playerID).ExecuteString()
	fmt.Println(matchUser)

	if err != nil {
		return false, err
	}

	isDealer := matchUser == "true"

	return isDealer, err
}

func (r *SupabaseMatchUsersRepository) GetAlivePlayers(matchID string) ([]models.MatchUsers, error) {
	var alivePlayers []models.MatchUsers

	matchUsers, _, err := r.DB.From("match_users").Select("*", "", false).Eq("match_id", matchID).ExecuteString()
	fmt.Println(matchUsers)

	return alivePlayers, err
}
