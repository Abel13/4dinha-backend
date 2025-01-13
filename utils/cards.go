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

func RemoveCards(cards []models.Deck, symbolsToRemove []string) []models.Deck {
	toRemove := make(map[string]bool)
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
) []map[string]interface{} {
	var distributedCards []map[string]interface{}

	for i := 0; i < cardsPerPlayer; i++ {
		for _, player := range players {
			if len(*shuffledCards) == 0 {
				break
			}

			cardID := (*shuffledCards)[0].ID
			*shuffledCards = (*shuffledCards)[1:]

			card := map[string]interface{}{
				"match_user":   player.ID,
				"round_number": roundNumber,
				"card_id":      cardID,
			}

			distributedCards = append(distributedCards, card)
		}
	}

	return distributedCards
}
