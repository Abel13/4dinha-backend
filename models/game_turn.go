package models

type GameTurn struct {
	TurnNumber int                  `json:"turn_number"`
	WinnerCard GetPlayerCardsResult `json:"winner_card"`
}
