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

type PlayerCardsRepository struct {
	SupabaseServiceRole string
	SupabaseURL         string
	SupabaseKey         string
}

func NewServicePlayerCardsRepository(supabaseServiceRole, supabaseURL, supabaseKey string) *PlayerCardsRepository {
	return &PlayerCardsRepository{
		SupabaseServiceRole: supabaseServiceRole,
		SupabaseURL:         supabaseURL,
		SupabaseKey:         supabaseKey,
	}
}

func NewPlayerCardsRepository() *PlayerCardsRepository {
	return &PlayerCardsRepository{}
}

func (r *PlayerCardsRepository) CreatePlayerCards(client *supabase.Client, playerCards []models.PlayerCardsInput) error {
	_, _, err := client.
		From("player_cards").
		Insert(playerCards, false, "", "minimal", "").
		Execute()

	if err != nil {
		return err
	}

	return nil
}

func (r *PlayerCardsRepository) GetPlayerCards(matchID, userID string, roundNumber int) ([]models.GetPlayerCardsResult, error) {
	body := models.GetPlayerCardsArgs{
		MatchID:     matchID,
		UserID:      userID,
		RoundNumber: roundNumber,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter corpo para JSON: %v", err)
	}

	url := fmt.Sprintf("%s/rest/v1/rpc/get_player_cards", r.SupabaseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição HTTP: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.SupabaseServiceRole))
	req.Header.Set("apikey", r.SupabaseKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("requisição falhou com status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result []models.GetPlayerCardsResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta JSON: %v", err)
	}

	return result, nil
}

func (r *PlayerCardsRepository) GetPlayerCardByID(client *supabase.Client, playerCardID string) (models.PlayerCards, error) {
	var playerCard models.PlayerCards

	_, err := client.
		From("player_cards").
		Select("*", "", false).
		Eq("id", playerCardID).
		Single().
		ExecuteTo(&playerCard)

	return playerCard, err
}

func (r *PlayerCardsRepository) Play(client *supabase.Client, playerCardID string) error {
	playedCard := models.PlayerCardsUpdate{
		Status: models.StatusOnTable,
	}

	_, _, err := client.
		From("player_cards").
		Update(playedCard, "minimal", "").
		Eq("id", playerCardID).
		Execute()

	if err != nil {
		return err
	}

	return nil
}
