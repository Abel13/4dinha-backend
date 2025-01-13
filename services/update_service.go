package services

import (
	"4dinha-backend/models"
	"4dinha-backend/repositories"
	"4dinha-backend/utils"
	"errors"
	"fmt"
	"strconv"

	"github.com/supabase-community/supabase-go"
)

type UpdateService struct {
	MatchRepo       repositories.MatchRepository
	MatchUserRepo   repositories.MatchUsersRepository
	PlayerCardsRepo repositories.PlayerCardsRepository
	RoundRepo       repositories.RoundRepository
	DeckRepo        repositories.DeckRepository
	ActionRepo      repositories.MatchActionRepository
	BetRepo         repositories.BetRepository
}

type GameUpdate struct {
	Match       models.Matches                `json:"match"`
	Players     []models.GamePlayer           `json:"players"`
	PlayerCards []models.GetPlayerCardsResult `json:"player_cards"`
	Round       models.GameRound              `json:"round"`
	Bets        []models.Bets                 `json:"bets"`
	Results     []models.PlayerResult         `json:"results"`
}

func NewUpdateService(
	matchRepo repositories.MatchRepository,
	matchUserRepo repositories.MatchUsersRepository,
	playerCardRepo repositories.PlayerCardsRepository,
	roundRepo repositories.RoundRepository,
	deckRepo repositories.DeckRepository,
	actionRepo repositories.MatchActionRepository,
	betRepo repositories.BetRepository,
) *UpdateService {
	return &UpdateService{
		MatchRepo:       matchRepo,
		MatchUserRepo:   matchUserRepo,
		PlayerCardsRepo: playerCardRepo,
		RoundRepo:       roundRepo,
		DeckRepo:        deckRepo,
		ActionRepo:      actionRepo,
		BetRepo:         betRepo,
	}
}

func FindMatchUserByUserID(matchUsers []models.MatchUsers, userID string) (models.MatchUsers, bool) {
	for _, user := range matchUsers {
		if user.UserID == userID {
			return user, true
		}
	}

	return models.MatchUsers{}, false
}

func GetResult(
	maxTurns int,
	allPlayerCards []models.GetPlayerCardsResult,
	trumpPower int,
	deck []models.Deck,
	players []models.MatchUsers,
	bets []models.Bets) []models.PlayerResult {
	playersResults := make(map[string]*models.PlayerResult)

	for _, player := range players {
		playersResults[player.UserID] = &models.PlayerResult{
			PlayerID: player.UserID,
			Lives:    player.Lives,
			Bets:     0,
			Wins:     0,
		}

		for _, bet := range bets {
			if bet.UserID == player.UserID {
				playersResults[player.UserID].Bets = bet.Bet
				break
			}
		}
	}

	for turn := 1; turn <= maxTurns; turn++ {
		var playedCards []models.Deck

		for _, playedCard := range allPlayerCards {
			if playedCard.Turn == turn {
				for _, card := range deck {
					if card.Symbol == playedCard.Symbol && card.Suit == playedCard.Suit {
						playedCards = append(playedCards, card)
						break
					}
				}
			}
		}

		for i := range playedCards {
			if playedCards[i].Power == trumpPower {
				playedCards[i].Power += 13
			}
		}

		var winningCard models.Deck
		for _, card := range playedCards {
			if winningCard.ID == "" || card.Power > winningCard.Power ||
				(card.Power == winningCard.Power && card.SuitPower > winningCard.SuitPower) {
				winningCard = card
			}
		}

		for _, playedCard := range allPlayerCards {
			if playedCard.Symbol == winningCard.Symbol && playedCard.Suit == winningCard.Suit {
				if result, exists := playersResults[playedCard.UserID]; exists {
					result.Wins++
					break
				}
			}
		}
	}

	var results []models.PlayerResult
	for _, result := range playersResults {
		results = append(results, *result)
	}

	return results
}

func FindNextAlivePlayer(matchUsers []models.MatchUsers, lastPlayerID string) (models.MatchUsers, bool) {
	lastPlayer, found := FindMatchUserByUserID(matchUsers, lastPlayerID)
	if !found {
		return models.MatchUsers{}, false
	}

	totalPlayers := len(matchUsers)
	currentSeat := *lastPlayer.TableSeat

	for {
		nextSeat := (currentSeat % totalPlayers) + 1

		for _, player := range matchUsers {
			if player.TableSeat != nil && *player.TableSeat == nextSeat {
				if player.Lives > 0 {
					return player, true
				}
				break
			}
		}

		currentSeat = nextSeat

		if currentSeat == *lastPlayer.TableSeat {
			break
		}
	}

	return models.MatchUsers{}, false
}

func (s *UpdateService) Update(client *supabase.Client, matchID, playerID string) (GameUpdate, error) {
	var gamePlayers []models.GamePlayer
	var gameUpdate GameUpdate
	var currentPlayerID string
	var results []models.PlayerResult

	match, err := s.MatchRepo.GetMatch(client, matchID)
	if err != nil {
		return gameUpdate, errors.New("error getting match")
	}
	stringRoundNumber := strconv.Itoa(match.RoundNumber)

	matchUsers, err := s.MatchUserRepo.GetAlivePlayers(client, matchID)
	if err != nil {
		return gameUpdate, fmt.Errorf("erro ao buscar jogadores: %v", err)
	}

	playerCards, err := s.PlayerCardsRepo.GetPlayerCards(matchID, playerID, match.RoundNumber)
	if err != nil {
		return gameUpdate, fmt.Errorf("erro ao buscar cartas do jogador: %v", err)
	}

	gameBets, err := s.BetRepo.GetRoundBets(client, matchID, stringRoundNumber)

	round, err := s.RoundRepo.CurrentRound(client, stringRoundNumber, matchID)
	if err == nil {
		card := s.DeckRepo.GetCard(client, round.Trump)
		gameRound := models.GameRound{
			RoundNumber: round.RoundNumber,
			Status:      round.Status,
			TrumpSymbol: card.Symbol,
			TrumpSuit:   card.Suit,
		}

		gameUpdate.Round = gameRound

		if gameRound.Status == models.StatusFinished {
			trumpIndicator := s.DeckRepo.GetCard(client, round.Trump)
			deck := s.DeckRepo.GetAllCards(client)
			results = GetResult(utils.CalculateGroup(match.RoundNumber), playerCards, (trumpIndicator.Power%13)+1, deck, matchUsers, gameBets)
		} else {
			lastAction, err := s.ActionRepo.GetLastAction(client, matchID)
			if err != nil {
				return gameUpdate, fmt.Errorf("erro ao buscar ultima acao: %v", err)
			}

			lastPlayer, found := FindMatchUserByUserID(matchUsers, lastAction.UserID)
			nextPlayer, found := FindNextAlivePlayer(matchUsers, lastPlayer.UserID)

			if found {
				currentPlayerID = nextPlayer.UserID
			}
			//if found {
			//	nextSeat := (*lastPlayer.TableSeat % len(matchUsers)) + 1
			//	nextSeatString := strconv.Itoa(nextSeat)
			//	nextPlayer, err := s.MatchUserRepo.GetPlayerBySeat(client, matchID, nextSeatString)
			//	if err != nil {
			//		return gameUpdate, fmt.Errorf("erro ao buscar proximo jogador: %v", err)
			//	}
			//
			//	currentPlayerID = nextPlayer.UserID
			//}
		}
	}

	for _, matchUser := range matchUsers {
		gamePlayer := models.GamePlayer{
			ID:        matchUser.ID,
			UserID:    matchUser.UserID,
			Lives:     matchUser.Lives,
			TableSeat: matchUser.TableSeat,
			Dealer:    matchUser.Dealer,
			Current:   matchUser.UserID == currentPlayerID,
		}

		gamePlayers = append(gamePlayers, gamePlayer)
	}

	gameUpdate.Match = match
	gameUpdate.Players = gamePlayers
	gameUpdate.PlayerCards = playerCards
	gameUpdate.Bets = gameBets
	gameUpdate.Results = results

	return gameUpdate, nil
}
