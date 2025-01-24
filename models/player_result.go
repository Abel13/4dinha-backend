package models

type PlayerResult struct {
	PlayerID string `json:"user_id"`
	Lives    int    `json:"lives"`
	Bets     int    `json:"bets"`
	Wins     int    `json:"wins"`
}

type RoundResult struct {
	PlayersResult []PlayerResult
	LastWinnerID  string
}
