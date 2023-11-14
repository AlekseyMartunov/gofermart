package users

import "time"

type User struct {
	ID        string
	UUID      string
	Login     string
	Password  string
	Withdrawn float64
	Bonuses   float64
}

type HistoryElement struct {
	Order        string    `json:"order"`
	Amount       float64   `json:"sum"`
	WriteOffTime time.Time `json:"processed_at"`
}
