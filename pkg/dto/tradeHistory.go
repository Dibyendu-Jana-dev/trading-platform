package dto

import "time"

type TradeHistoryDTO struct {
	ID        string    `json:"id"`
	Symbol    string    `json:"symbol"`
	Volume    float64   `json:"volume"`
	Price     float64   `json:"price"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}