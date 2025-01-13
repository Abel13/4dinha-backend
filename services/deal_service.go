package services

import (
	"4dinha-backend/repositories"
	"errors"
	"github.com/supabase-community/supabase-go"
)

type DealService struct {
	MatchRepo     repositories.MatchRepository
	MatchUserRepo repositories.MatchUsersRepository
	DeckRepo      repositories.DeckRepository
}

func NewDealService(matchRepo repositories.MatchRepository, matchUserRepo repositories.MatchUsersRepository, deckRepo repositories.DeckRepository) *DealService {
	return &DealService{
		MatchRepo:     matchRepo,
		MatchUserRepo: matchUserRepo,
		DeckRepo:      deckRepo,
	}
}

func (s *DealService) DealCards(client *supabase.Client, userID, matchID string) error {

	isDealer, err := s.MatchUserRepo.IsDealer(client, matchID, userID)
	if err != nil || !isDealer {
		return errors.New("user is not the dealer")
	}

	return nil
}
