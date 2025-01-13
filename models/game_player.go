package models

type GamePlayer struct {
	ID         string `json:"id"` // User ID
	UserID     string `json:"user_id"`
	Dealer     bool   `json:"dealer"`
	Current    bool   `json:"current"`
	LastPlayer bool   `json:"last_player"`
	Lives      int    `json:"lives"`
	TableSeat  *int   `json:"table_seat"`
}
