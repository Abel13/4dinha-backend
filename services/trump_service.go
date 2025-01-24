package services

import (
	"4dinha-backend/models"
	"4dinha-backend/repositories"
	"errors"
	"strconv"

	"github.com/supabase-community/supabase-go"
)

type TrumpService struct {
	MatchRepo repositories.MatchRepository
	RoundRepo repositories.RoundRepository
	DeckRepo  repositories.DeckRepository
}

func NewTrumpService(
	matchRepo repositories.MatchRepository,
	roundRepo repositories.RoundRepository,
	deckRepo repositories.DeckRepository) *TrumpService {
	return &TrumpService{
		MatchRepo: matchRepo,
		RoundRepo: roundRepo,
		DeckRepo:  deckRepo,
	}
}

func GetTrumpPower(trumpDesignator models.Deck) int {
	var trumpPower int

	// REMOVE 8s e 9s
	for trumpPower == 0 || (trumpPower >= 5 && trumpPower <= 7) {
		if trumpPower == 0 {
			trumpPower = trumpDesignator.Power
		}

		trumpPower = (trumpPower % 13) + 1
	}

	return trumpPower
}

func (s *TrumpService) GetTrumps(client *supabase.Client, matchID string) ([]models.Deck, error) {
	match, err := s.MatchRepo.GetMatch(client, matchID)
	if err != nil {
		return nil, errors.New("error getting match")
	}
	stringRoundNumber := strconv.Itoa(match.RoundNumber)
	round, err := s.RoundRepo.CurrentRound(client, stringRoundNumber, matchID)
	if err != nil {
		return nil, errors.New("error getting trumps")
	}

	trumpDesignator := s.DeckRepo.GetCard(client, round.Trump)
	trumpPower := GetTrumpPower(trumpDesignator)
	trumps, err := s.RoundRepo.GetTrumpsByPower(client, strconv.Itoa(trumpPower))
	if err != nil {
		return nil, errors.New("error getting trumps")
	}

	return trumps, err
}
