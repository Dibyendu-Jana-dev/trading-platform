package domain

import (
	"context"
	"time"

	"github.com/dibyendu/trading_platform/lib/errs"
)

type PositionRepository interface {
	// PlaceOrder(ctx context.Context, req Order) (*Order, *errs.AppError)
	GetUserPositions(ctx context.Context, userID string) ([]*Position, *errs.AppError)
}

type Position struct {
	UserID    string    `json:"user_id" bson:"user_id"`
	Symbol    string    `json:"symbol" bson:"symbol"`
	Quantity  float64   `json:"quantity" bson:"quantity"`
	AvgPrice  float64   `json:"avg_price" bson:"avg_price"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}