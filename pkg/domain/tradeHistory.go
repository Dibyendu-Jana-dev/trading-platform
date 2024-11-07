package domain

import (
	"context"
	"time"

	"github.com/dibyendu/trading_platform/lib/errs"
)

type TradingHistoryRepository interface {
	GetTradeHistoryByUserID(ctx context.Context, userId string) ([]*TradeHistory, *errs.AppError)
}

type TradeHistory struct {
	ID        string    `json:"id" bson:"_id"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Symbol    string    `json:"symbol" bson:"symbol"`
	Volume    float64   `json:"volume" bson:"volume"`
	Price     float64   `json:"price" bson:"price"`
	Type      string    `json:"type" bson:"type"` // "buy" or "sell"
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}