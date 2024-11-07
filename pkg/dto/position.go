package dto

import "time"

type PositionDTO struct {
	Symbol    string    `json:"symbol"`
	Quantity  float64   `json:"quantity"`
	AvgPrice  float64   `json:"avg_price"`
	Timestamp time.Time `json:"timestamp"`
}