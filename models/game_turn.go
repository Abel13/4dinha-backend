package models

type GameTurn struct {
	TurnNumber int  `json:"turn_number"`
	WinnerCard Deck `json:"winner_card"`
}
