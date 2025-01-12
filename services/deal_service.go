package services

import (
	"4dinha-backend/repositories"
	"errors"
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

func (s *DealService) DealCards(userID, matchID string) error {

	isDealer, err := s.MatchUserRepo.IsDealer(matchID, userID)
	if err != nil || !isDealer {
		return errors.New("user is not the dealer")
	}

	// Busque todos os jogadores vivos
	//players, err := s.MatchUserRepo.GetAlivePlayers(matchID)
	//if err != nil {
	//	return err
	//}

	// Busque todas as cartas
	//cards, err := s.DeckRepo.GetAllCards()
	//if err != nil {
	//	return err
	//}

	// Distribua as cartas e insira em player_cards
	// (lógica completa de distribuição omitida por brevidade)

	return nil
}
