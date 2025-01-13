package repositories

import (
	"4dinha-backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/supabase-community/supabase-go"
	"io"
	"net/http"
)

type MatchRepository struct {
	SupabaseServiceRole string
	SupabaseURL         string
	SupabaseKey         string
}

func NewServiceMatchRepository(supabaseServiceRole, supabaseURL, supabaseKey string) *MatchRepository {
	return &MatchRepository{
		SupabaseServiceRole: supabaseServiceRole,
		SupabaseURL:         supabaseURL,
		SupabaseKey:         supabaseKey,
	}
}

func NewMatchRepository() *MatchRepository {
	return &MatchRepository{}
}

func (r *MatchRepository) GetMatch(client *supabase.Client, matchID string) (models.Matches, error) {
	var match models.Matches

	_, err := client.
		From("matches").
		Select("*", "", false).
		Eq("id", matchID).
		Eq("status", string(models.StatusStarted)).
		Single().
		ExecuteTo(&match)

	return match, err
}

func (r *MatchRepository) UpdateRoundNumber(matchID string, newRoundNumber int) error {
	body := models.UpdateRoundNumberArgs{
		MatchID:        matchID,
		NewRoundNumber: newRoundNumber,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("erro ao converter corpo para JSON: %v", err)
	}

	url := fmt.Sprintf("%s/rest/v1/rpc/update_round_number", r.SupabaseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("erro ao criar requisição HTTP: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.SupabaseServiceRole))
	req.Header.Set("apikey", r.SupabaseKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erro ao executar requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("requisição falhou com status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
