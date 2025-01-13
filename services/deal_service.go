package services

import (
	"4dinha-backend/repositories"
	"4dinha-backend/utils"
	"errors"
	"fmt"
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

	players, err := s.MatchUserRepo.GetAlivePlayers(client, matchID)
	if err != nil {
		return errors.New("error getting players")
	}
	match, err := s.MatchRepo.GetMatch(client, matchID)
	if err != nil {
		return errors.New("error getting match")
	}

	deck, err := s.DeckRepo.GetAllCards(client)
	unusedSymbols := []string{"8", "9", "10"}
	gameDeck := utils.RemoveCards(deck, unusedSymbols)
	roundNumber := match.RoundNumber

	shuffledCards := utils.Shuffle(gameDeck)

	cardsQuantity := utils.CalculateGroup(roundNumber)

	playerCards := utils.DistributeCards(players, roundNumber, &shuffledCards, cardsQuantity)

	fmt.Println(len(playerCards), len(shuffledCards))

	return nil
}
