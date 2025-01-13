package services

import (
	"4dinha-backend/models"
	"4dinha-backend/repositories"
	"4dinha-backend/utils"
	"errors"
	"github.com/supabase-community/supabase-go"
)

type DealService struct {
	MatchRepo      repositories.MatchRepository
	MatchUserRepo  repositories.MatchUsersRepository
	DeckRepo       repositories.DeckRepository
	PlayerCardRepo repositories.PlayerCardsRepository
	RoundRepo      repositories.RoundRepo
}

func NewDealService(
	matchRepo repositories.MatchRepository,
	matchUserRepo repositories.MatchUsersRepository,
	deckRepo repositories.DeckRepository,
	playerCardRepo repositories.PlayerCardsRepository,
	roundRepo repositories.RoundRepo) *DealService {
	return &DealService{
		MatchRepo:      matchRepo,
		MatchUserRepo:  matchUserRepo,
		DeckRepo:       deckRepo,
		PlayerCardRepo: playerCardRepo,
		RoundRepo:      roundRepo,
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
	roundNumber := match.RoundNumber

	err = s.RoundRepo.CreateRound(client, match.ID, roundNumber)
	if err != nil {
		return errors.New("error creating round")
	}

	deck, err := s.DeckRepo.GetAllCards(client)
	unusedSymbols := []models.CardSymbol{models.Symbol8, models.Symbol9, models.Symbol10}
	gameDeck := utils.RemoveCards(deck, unusedSymbols)

	shuffledCards := utils.Shuffle(gameDeck)

	cardsQuantity := utils.CalculateGroup(roundNumber)

	playerCards := utils.DistributeCards(players, roundNumber, &shuffledCards, cardsQuantity)

	err = s.PlayerCardRepo.CreatePlayerCards(client, playerCards)
	if err != nil {
		return errors.New("error dealing cards")
	}

	return nil
}
