package services

import (
	"4dinha-backend/models"
	"4dinha-backend/repositories"
	"4dinha-backend/utils"
	"errors"
	"strconv"

	"github.com/supabase-community/supabase-go"
)

type DealService struct {
	MatchRepo       repositories.MatchRepository
	MatchUserRepo   repositories.MatchUsersRepository
	MatchActionRepo repositories.MatchActionRepository
	DeckRepo        repositories.DeckRepository
	PlayerCardRepo  repositories.PlayerCardsRepository
	RoundRepo       repositories.RoundRepository
}

func NewDealService(
	matchRepo repositories.MatchRepository,
	matchUserRepo repositories.MatchUsersRepository,
	matchActionRepo repositories.MatchActionRepository,
	deckRepo repositories.DeckRepository,
	playerCardRepo repositories.PlayerCardsRepository,
	roundRepo repositories.RoundRepository) *DealService {
	return &DealService{
		MatchRepo:       matchRepo,
		MatchUserRepo:   matchUserRepo,
		DeckRepo:        deckRepo,
		PlayerCardRepo:  playerCardRepo,
		RoundRepo:       roundRepo,
		MatchActionRepo: matchActionRepo,
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
	stringRoundNumber := strconv.Itoa(match.RoundNumber)

	exists, _ := s.RoundRepo.CurrentRound(client, stringRoundNumber, matchID)

	if exists != nil {
		return errors.New("Round already exists")
	}

	deck := s.DeckRepo.GetAllCards(client)
	unusedSymbols := []models.CardSymbol{models.Symbol8, models.Symbol9, models.Symbol10}
	gameDeck := utils.RemoveCards(deck, unusedSymbols)

	shuffledCards := utils.Shuffle(gameDeck)

	cardsQuantity := utils.CalculateGroup(roundNumber)

	playerCards := utils.DistributeCards(players, roundNumber, &shuffledCards, cardsQuantity)

	err = s.PlayerCardRepo.CreatePlayerCards(client, playerCards)
	if err != nil {
		return errors.New("error dealing cards")
	}

	err = s.RoundRepo.CreateRound(client, match.ID, roundNumber, &shuffledCards)
	if err != nil {
		return errors.New("error creating round")
	}

	err = s.MatchActionRepo.RegisterAction(client, matchID, roundNumber, models.ActionDeal)
	if err != nil {
		return errors.New("error registering deal action")
	}

	return nil
}
