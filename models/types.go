package models

// Tables
type Deck struct {
	ID        string     `json:"id"`
	Power     int        `json:"power"`
	Suit      CardSuit   `json:"suit"`
	SuitPower int        `json:"suit_power"`
	Symbol    CardSymbol `json:"symbol"`
}

type MatchActions struct {
	ID          string  `json:"id"`
	Action      Actions `json:"action"`
	MatchID     string  `json:"match_id"`
	RoundNumber int     `json:"round_number"`
	UserID      string  `json:"user_id"`
	CreatedAt   string  `json:"created_at"`
}

type MatchActionsInput struct {
	Action      Actions `json:"action"`
	MatchID     string  `json:"match_id"`
	RoundNumber int     `json:"round_number"`
}

type MatchUsers struct {
	CreatedAt *string `json:"created_at"`
	Dealer    bool    `json:"dealer"`
	ID        string  `json:"id"`
	Lives     int     `json:"lives"`
	MatchID   string  `json:"match_id"`
	Ready     bool    `json:"ready"`
	TableSeat *int    `json:"table_seat"`
	UserID    string  `json:"user_id"`
}

type Matches struct {
	CreatedAt   *string      `json:"created_at"`
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	RoundNumber int          `json:"round_number"`
	Status      *MatchStatus `json:"status"`
	UserID      string       `json:"user_id"`
}

type PlayerCards struct {
	ID          string     `json:"id"`
	CardID      string     `json:"card_id"`
	CreatedAt   string     `json:"created_at"`
	MatchID     string     `json:"match_id"`
	RoundNumber int        `json:"round_number"`
	Status      HandStatus `json:"status"`
	UserID      string     `json:"user_id"`
	Turn        *int       `json:"turn"`
}

type PlayerCardsInput struct {
	CardID      string     `json:"card_id"`
	MatchID     string     `json:"match_id"`
	RoundNumber int        `json:"round_number"`
	Status      HandStatus `json:"status"`
	UserID      string     `json:"user_id"`
	Turn        *int       `json:"turn"`
}

type PlayerCardsUpdate struct {
	Status HandStatus `json:"status"`
	Turn   *int       `json:"turn"`
}

type Rounds struct {
	MatchID     string      `json:"match_id"`
	RoundNumber int         `json:"round_number"`
	Status      RoundStatus `json:"status"`
	Trump       string      `json:"trump"`
}

type Bets struct {
	Bet         int     `json:"bet"`
	MatchID     string  `json:"match_id"`
	RoundNumber int     `json:"round_number"`
	UserID      string  `json:"user_id"`
	CreatedAt   *string `json:"created_at"`
}

type BetsInput struct {
	Bet         int    `json:"bet"`
	MatchID     string `json:"match_id"`
	RoundNumber int    `json:"round_number"`
}

// Enums
type Actions string

const (
	ActionDeal         Actions = "deal"
	ActionBet          Actions = "bet"
	ActionChangeStatus Actions = "change_status"
	ActionPlay         Actions = "play"
	ActionRoundStart   Actions = "round_start"
)

type CardSuit string

const (
	SuitClubs    CardSuit = "♣️"
	SuitHearts   CardSuit = "♥️"
	SuitSpades   CardSuit = "♠️"
	SuitDiamonds CardSuit = "♦️"
)

type CardSymbol string

const (
	SymbolA  CardSymbol = "A"
	Symbol2  CardSymbol = "2"
	Symbol3  CardSymbol = "3"
	Symbol4  CardSymbol = "4"
	Symbol5  CardSymbol = "5"
	Symbol6  CardSymbol = "6"
	Symbol7  CardSymbol = "7"
	Symbol8  CardSymbol = "8"
	Symbol9  CardSymbol = "9"
	Symbol10 CardSymbol = "10"
	SymbolQ  CardSymbol = "Q"
	SymbolJ  CardSymbol = "J"
	SymbolK  CardSymbol = "K"
)

type HandStatus string

const (
	StatusOnHand HandStatus = "on hand"
	StatusPlayed HandStatus = "played"
)

type MatchStatus string

const (
	StatusCreated MatchStatus = "created"
	StatusStarted MatchStatus = "started"
	StatusEnd     MatchStatus = "end"
)

type RoundStatus string

const (
	StatusDealing  RoundStatus = "dealing"
	StatusBetting  RoundStatus = "betting"
	StatusPlaying  RoundStatus = "playing"
	StatusFinished RoundStatus = "finished"
)

// Stored Procedures and Functions
type GetUserEmailArgs struct {
	UserID string `json:"user_id"`
}

type GetUserEmailResult struct {
	Email string `json:"email"`
}

type UpdateMatchStatusToStartedArgs struct {
	MatchID string `json:"_match_id"`
}

type UpdateReadyStatusArgs struct {
	MatchID string `json:"_match_id"`
	IsReady bool   `json:"_is_ready"`
}

type UpdateDealerArgs struct {
	MatchID   string `json:"_match_id"`
	TableSeat int    `json:"_table_seat"`
}

type UpdatePlayerLivesArgs struct {
	MatchID  string `json:"_match_id"`
	UserID   string `json:"_user_id"`
	NewLives int    `json:"_new_lives"`
}

type UpdatePlayerLivesResult struct {
}

type UpdateRoundNumberArgs struct {
	MatchID        string `json:"_match_id"`         // ID do match a ser atualizado
	NewRoundNumber int    `json:"_new_round_number"` // Novo número da rodada
}

type GetPlayerCardsArgs struct {
	MatchID     string `json:"_match_id"`
	UserID      string `json:"_user_id"`
	RoundNumber int    `json:"_round_number"`
}

type GetPlayerCardsResult struct {
	ID     string     `json:"id"`
	UserID string     `json:"user_id"`
	Symbol CardSymbol `json:"symbol"`
	Suit   CardSuit   `json:"suit"`
	Status HandStatus `json:"status"`
	Turn   int        `json:"turn"`
}
