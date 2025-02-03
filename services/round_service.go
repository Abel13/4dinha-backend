package services

import (
	"4dinha-backend/models"
	"4dinha-backend/repositories"
	"4dinha-backend/utils"
	"fmt"
	"github.com/supabase-community/supabase-go"
	"math"
	"sort"
	"strconv"
)

type RoundService struct {
	MatchRepo       repositories.MatchRepository
	RoundRepo       repositories.RoundRepository
	PlayerCardsRepo repositories.PlayerCardsRepository
	DeckRepo        repositories.DeckRepository
	MatchUsersRepo  repositories.MatchUsersRepository
	BetRepo         repositories.BetRepository
}

func NewRoundService(
	matchRepo repositories.MatchRepository,
	roundRepo repositories.RoundRepository,
	playerCardsRepo repositories.PlayerCardsRepository,
	deckRepo repositories.DeckRepository,
	mathUsersRepo repositories.MatchUsersRepository,
	betRepo repositories.BetRepository) *RoundService {
	return &RoundService{
		MatchRepo:       matchRepo,
		RoundRepo:       roundRepo,
		PlayerCardsRepo: playerCardsRepo,
		DeckRepo:        deckRepo,
		MatchUsersRepo:  mathUsersRepo,
		BetRepo:         betRepo,
	}
}

func GetNextMatchUser(users []models.MatchUsers, currentTableSeat int) *models.MatchUsers {
	sort.Slice(users, func(i, j int) bool {
		return *users[i].TableSeat < *users[j].TableSeat
	})

	for _, user := range users {
		if *user.TableSeat > currentTableSeat {
			return &user
		}
	}

	if len(users) > 0 {
		return &users[0]
	}

	return nil
}

func FindMatchUserByID(users []models.MatchUsers, userID string) *models.MatchUsers {
	for _, user := range users {
		if user.UserID == userID {
			return &user
		}
	}
	return nil
}

func (s *RoundService) FinishRound(client *supabase.Client, matchID string, playerID string) error {
	matchUsers, _ := s.MatchUsersRepo.GetAlivePlayers(client, matchID)
	user := FindMatchUserByID(matchUsers, playerID)

	if user == nil || user.Dealer == false {
		return fmt.Errorf("user is not dealer")
	}

	match, _ := s.MatchRepo.GetMatch(client, matchID)
	stringRoundNumber := strconv.Itoa(match.RoundNumber)
	round, _ := s.RoundRepo.CurrentRound(client, stringRoundNumber, matchID)
	playerCards, _ := s.PlayerCardsRepo.GetPlayerCards(matchID, playerID, match.RoundNumber)
	trumpIndicator := s.DeckRepo.GetCard(client, round.Trump)
	trumpPower := GetTrumpPower(trumpIndicator)
	deck := s.DeckRepo.GetAllCards(client)

	gameBets, _ := s.BetRepo.GetRoundBets(client, matchID, stringRoundNumber)

	results := GetResult(utils.CalculateGroup(match.RoundNumber), playerCards, trumpPower, deck, matchUsers, gameBets)
	if len(results.PlayersResult) > 0 {
		for _, result := range results.PlayersResult {
			lostLife := int(math.Abs(float64(result.Bets - result.Wins)))
			if lostLife == 0 {
				continue
			}

			life := result.Lives - lostLife

			err := s.MatchUsersRepo.UpdateLives(matchID, result.PlayerID, life)
			if err != nil {
				return err
			}
		}
	}

	// Update match round_number
	err := s.MatchRepo.UpdateRoundNumber(matchID, match.RoundNumber+1)
	if err != nil {
		return err
	}

	// getAlivePlayers
	matchUsers, _ = s.MatchUsersRepo.GetAlivePlayers(client, matchID)

	// choose next Dealer
	nextUser := GetNextMatchUser(matchUsers, *user.TableSeat)
	err = s.MatchUsersRepo.UpdateDealer(matchID, *nextUser.TableSeat)
	if err != nil {
		return err
	}

	return nil
}
