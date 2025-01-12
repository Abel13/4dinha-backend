package repositories

import (
	"4dinha-backend/models"
	"fmt"
	"github.com/supabase-community/supabase-go"
)

type MatchRepository interface {
	GetMatch() ([]models.Matches, error)
}

type SupabaseMatchRepository struct {
	DB *supabase.Client
}

func (r *SupabaseMatchRepository) GetMatch() ([]models.Matches, error) {
	var matchList []models.Matches

	matches, _, err := r.DB.From("matches").Select("*", "", false).Execute()

	fmt.Println(matches)
	return matchList, err
}
