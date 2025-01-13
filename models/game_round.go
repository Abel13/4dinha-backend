package models

type GameRound struct {
	RoundNumber int         `json:"round_number"`
	Status      RoundStatus `json:"status"`
	TrumpSymbol CardSymbol  `json:"trump_symbol"`
	TrumpSuit   CardSuit    `json:"trump_suit"`
}
