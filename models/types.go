package models

type Deck struct {
	ID        string     `json:"id"`
	Power     int        `json:"power"`
	Suit      CardSuit   `json:"suit"`
	SuitPower int        `json:"suit_power"`
	Symbol    CardSymbol `json:"symbol"`
}

type MatchActions struct {
	Action      Actions `json:"action"`
	CreatedAt   string  `json:"created_at"`
	ID          int     `json:"id"`
	MatchID     string  `json:"match_id"`
	RoundNumber int     `json:"round_number"`
}

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

type Matches struct {
	CreatedAt   *string      `json:"created_at"`
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	RoundNumber int          `json:"round_number"`
	Status      *MatchStatus `json:"status"`
	UserID      string       `json:"user_id"`
}

type PlayerCards struct {
	CardID      *string    `json:"card_id"`
	MatchUser   string     `json:"match_user"`
	RoundNumber int        `json:"round_number"`
	Status      HandStatus `json:"status"`
}

type Rounds struct {
	MatchID     string      `json:"match_id"`
	RoundNumber int         `json:"round_number"`
	Status      RoundStatus `json:"status"`
}

// Enums
type Actions string

const (
	ActionDeal Actions = "deal"
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
	StatusEnd     MatchStatus = "finished"
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

type UpdateMatchStatusToStartedArgs struct {
	MatchID string `json:"_match_id"`
}

type UpdateReadyStatusArgs struct {
	MatchID string `json:"_match_id"`
	IsReady bool   `json:"_is_ready"`
}
