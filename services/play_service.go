package services

import (
	"4dinha-backend/models"
	"4dinha-backend/repositories"
	"errors"
	"github.com/supabase-community/supabase-go"
	"strconv"
)

type PlayService struct {
	PlayerCardRepo repositories.PlayerCardsRepository
	RoundRepo      repositories.RoundRepository
}

func NewPlayService(playerCardRepo repositories.PlayerCardsRepository, roundRepo repositories.RoundRepository) *PlayService {
	return &PlayService{
		PlayerCardRepo: playerCardRepo,
		RoundRepo:      roundRepo,
	}
}

func (s *PlayService) Play(client *supabase.Client, playerCardID string) error {
	playerCard, err := s.PlayerCardRepo.GetPlayerCardByID(client, playerCardID)
	if err != nil {
		return err
	}

	if playerCard.Status == models.StatusPlayed {
		return errors.New("card already played")
	}

	stringRoundNumber := strconv.Itoa(playerCard.RoundNumber)
	matchID := playerCard.MatchID

	round, err := s.RoundRepo.CurrentRound(client, stringRoundNumber, matchID)
	if err != nil {
		return err
	}
	if round.Status == models.StatusPlaying {
		s.PlayerCardRepo.Play(client, playerCardID)
	}

	return nil
}
