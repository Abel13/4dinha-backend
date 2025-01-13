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

type MatchUsersRepository struct {
	SupabaseServiceRole string
	SupabaseURL         string
	SupabaseKey         string
}

func NewServiceMatchUsersRepository(supabaseServiceRole, supabaseURL, supabaseKey string) *MatchUsersRepository {
	return &MatchUsersRepository{
		SupabaseServiceRole: supabaseServiceRole,
		SupabaseURL:         supabaseURL,
		SupabaseKey:         supabaseKey,
	}
}

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

	_, err := client.From("match_users").Select("*", "", false).Eq("match_id", matchID).Gt("lives", "0").ExecuteTo(&alivePlayers)

	return alivePlayers, err
}

func (r *MatchUsersRepository) GetPlayerBySeat(client *supabase.Client, matchID, seat string) (models.MatchUsers, error) {
	var playerSeat models.MatchUsers

	_, err := client.
		From("match_users").
		Select("*", "", false).
		Eq("match_id", matchID).
		Eq("table_seat", seat).
		Single().
		ExecuteTo(&playerSeat)

	return playerSeat, err
}

func (r *MatchUsersRepository) UpdateLives(matchID, playerID string, live int) error {
	body := models.UpdatePlayerLivesArgs{
		MatchID:  matchID,
		UserID:   playerID,
		NewLives: live,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("erro ao converter corpo para JSON: %v", err)
	}

	url := fmt.Sprintf("%s/rest/v1/rpc/update_player_lives", r.SupabaseURL)
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

func (r *MatchUsersRepository) UpdateDealer(matchID string, tableSeat int) error {
	body := models.UpdateDealerArgs{
		MatchID:   matchID,
		TableSeat: tableSeat,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("erro ao converter corpo para JSON: %v", err)
	}

	url := fmt.Sprintf("%s/rest/v1/rpc/update_dealer", r.SupabaseURL)
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
