package models

// User representa o User do JWT
type User struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

// Deck representa a tabela `deck`
type Deck struct {
	ID        string `json:"id"`
	Power     int    `json:"power"`
	Suit      string `json:"suit"` // card_suit enum
	SuitPower int    `json:"suit_power"`
	Symbol    string `json:"symbol"` // card_symbol enum
}

// MatchActions representa a tabela `match_actions`
type MatchActions struct {
	Action      string `json:"action"` // actions enum
	CreatedAt   string `json:"created_at"`
	ID          int    `json:"id"`
	MatchID     string `json:"match_id"`
	RoundNumber int    `json:"round_number"`
}

// MatchUsers representa a tabela `match_users`
type MatchUsers struct {
	CreatedAt string `json:"created_at"`
	Dealer    bool   `json:"dealer"`
	ID        string `json:"id"`
	Lives     int    `json:"lives"`
	MatchID   string `json:"match_id"`
	Ready     bool   `json:"ready"`
	TableSit  *int   `json:"table_sit"`
	UserID    string `json:"user_id"`
}

// Matches representa a tabela `matches`
type Matches struct {
	CreatedAt   *string `json:"created_at"`
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	RoundNumber int     `json:"round_number"`
	Status      *string `json:"status"` // match_status enum
	UserID      string  `json:"user_id"`
}

// PlayerCards representa a tabela `player_cards`
type PlayerCards struct {
	CardID      *string `json:"card_id"`
	CreatedAt   string  `json:"created_at"`
	MatchUser   string  `json:"match_user"`
	RoundNumber int     `json:"round_number"`
}

// RoundOne representa a tabela `round_one`
type RoundOne struct {
	CreatedAt   string `json:"created_at"`
	ID          int    `json:"id"`
	RoundNumber int    `json:"round_number"`
	Status      string `json:"status"` // round_status enum
}

// Enums
const (
	// Actions enum
	ActionDeal = "deal"

	// CardSuit enum
	CardSuitClubs    = "♣️"
	CardSuitHearts   = "♥️"
	CardSuitSpades   = "♠️"
	CardSuitDiamonds = "♦️"

	// CardSymbol enum
	CardSymbolA  = "A"
	CardSymbol2  = "2"
	CardSymbol3  = "3"
	CardSymbol4  = "4"
	CardSymbol5  = "5"
	CardSymbol6  = "6"
	CardSymbol7  = "7"
	CardSymbol8  = "8"
	CardSymbol9  = "9"
	CardSymbol10 = "10"
	CardSymbolQ  = "Q"
	CardSymbolJ  = "J"
	CardSymbolK  = "K"

	// MatchStatus enum
	MatchStatusCreated  = "created"
	MatchStatusStarted  = "started"
	MatchStatusFinished = "finished"

	// RoundStatus enum
	RoundStatusDealing  = "dealing"
	RoundStatusBetting  = "betting"
	RoundStatusPlaying  = "playing"
	RoundStatusFinished = "finished"
)

// Functions

// GetUserEmailArgs representa os argumentos da função `get_user_email`
type GetUserEmailArgs struct {
	UserID string `json:"user_id"`
}

// GetUserEmailReturns representa os retornos da função `get_user_email`
type GetUserEmailReturns struct {
	Email string `json:"email"`
}

// UpdateMatchStatusToStartedArgs representa os argumentos da função `update_match_status_to_started`
type UpdateMatchStatusToStartedArgs struct {
	MatchID string `json:"_match_id"`
}

// UpdateReadyStatusArgs representa os argumentos da função `update_ready_status`
type UpdateReadyStatusArgs struct {
	MatchID string `json:"_match_id"`
	IsReady bool   `json:"_is_ready"`
}
