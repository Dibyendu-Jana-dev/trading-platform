package domain

import (
	"context"

	"github.com/dibyendu/trading_platform/lib/errs"
	"github.com/dibyendu/trading_platform/pkg/dto"
)

type MarketDataRepository interface {
	GetMarketData(ctx context.Context, symbol string) (*MarketData, *errs.AppError)
}

type MarketData struct {
	Symbol    string
	Price     string
	HighPrice string
	LowPrice  string
	Volume    string
	Timestamp int64
}

func (marketData MarketData) ToDto() *dto.MarketDataResponse{
	return &dto.MarketDataResponse{
		Symbol:    marketData.Symbol,
        Price:     marketData.Price,
        HighPrice: marketData.HighPrice,
        LowPrice:  marketData.LowPrice,
        Volume:    marketData.Volume,
        Timestamp: marketData.Timestamp,
	}
}
