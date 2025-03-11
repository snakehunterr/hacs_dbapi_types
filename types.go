package types

import (
	"time"
)

type Client struct {
	ID      uint   `json:"client_id"`
	Name    string `json:"client_name"`
	IsAdmin bool   `json:"is_admin"`
}

type Expense struct {
	ID     uint      `json:"expense_id"`
	Date   time.Time `json:"expense_date"`
	Amount float64   `json:"expense_amount"`
}

type Room struct {
	ID          uint    `json:"room_id"`
	ClientID    uint    `json:"client_id"`
	PeopleCount uint8   `json:"room_people_count"`
	Area        float64 `json:"room_area"`
}

type Payment struct {
	ID       uint      `json:"payment_id"`
	ClientID uint      `json:"client_id"`
	RoomID   uint      `json:"room_id"`
	Date     time.Time `json:"payment_date"`
	Amount   float64   `json:"payment_amount"`
}

type APIErrorCode byte

const (
	ErrMissingParam APIErrorCode = iota
	ErrIncorrectParam
	ErrSQLNoRows
	ErrSQLInternalError
)

type APIError struct {
	Code    APIErrorCode `json:"code"`
	Message string       `json:"message"`
}
