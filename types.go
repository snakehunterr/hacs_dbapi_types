package types

import (
	errors "github.com/snakehunterr/hacs_dbapi_types/errors"
	"time"
)

type Client struct {
	ID         uint      `json:"client_id"`
	Name       string    `json:"client_name"`
	IsAdmin    bool      `json:"is_admin"`
	LastEdited time.Time `json:"last_edited"`
}

type Expense struct {
	ID         uint      `json:"expense_id"`
	Date       time.Time `json:"expense_date"`
	Amount     float64   `json:"expense_amount"`
	LastEdited time.Time `json:"last_edited"`
}

type Room struct {
	ID          uint      `json:"room_id"`
	ClientID    uint      `json:"client_id"`
	PeopleCount uint8     `json:"room_people_count"`
	Area        float64   `json:"room_area"`
	LastEdited  time.Time `json:"last_edited"`
}

type Payment struct {
	ID         uint      `json:"payment_id"`
	ClientID   uint      `json:"client_id"`
	RoomID     uint      `json:"room_id"`
	Date       time.Time `json:"payment_date"`
	Amount     float64   `json:"payment_amount"`
	LastEdited time.Time `json:"last_edited"`
}

type APIResponse struct {
	Error   *errors.APIError `json:"error,omitempty"`
	Message string           `json:"message,omitempty"`
}
