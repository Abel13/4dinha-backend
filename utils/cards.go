package utils

import (
	"4dinha-backend/models"
	"math/rand"
)

func Shuffle(cards []models.Deck) []models.Deck {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func RemoveCards(cards []models.Deck, symbolsToRemove []models.CardSymbol) []models.Deck {
	toRemove := make(map[models.CardSymbol]bool)
	for _, symbol := range symbolsToRemove {
		toRemove[symbol] = true
	}

	var filteredDeck []models.Deck
	for _, card := range cards {
		if !toRemove[card.Symbol] {
			filteredDeck = append(filteredDeck, card)
		}
	}

	return filteredDeck
}

func CalculateGroup(roundNumber int) int {
	lastDigit := roundNumber % 10

	if lastDigit == 0 {
		return 2
	}
	if lastDigit <= 6 {
		return lastDigit
	}
	return 12 - lastDigit
}

func DistributeCards(
	players []models.MatchUsers,
	roundNumber int,
	shuffledCards *[]models.Deck,
	cardsPerPlayer int,
) []models.PlayerCards {
	var distributedCards []models.PlayerCards

	for i := 0; i < cardsPerPlayer; i++ {
		for _, player := range players {
			if len(*shuffledCards) == 0 {
				break
			}

			cardID := (*shuffledCards)[0].ID
			*shuffledCards = (*shuffledCards)[1:]

			card := models.PlayerCards{
				CardID:      &cardID,
				MatchUser:   player.ID,
				RoundNumber: roundNumber,
				Status:      models.StatusOnHand,
			}

			distributedCards = append(distributedCards, card)
		}
	}

	return distributedCards
}
