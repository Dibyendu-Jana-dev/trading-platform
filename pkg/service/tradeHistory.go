package service

import (
	"context"

	"github.com/dibyendu/trading_platform/lib/errs"
	"github.com/dibyendu/trading_platform/pkg/domain"
	"github.com/dibyendu/trading_platform/pkg/dto"
)

type DefaultTradingHistoryService struct {
	repo domain.TradingHistoryRepository
}

func NewTradingHistoryService(repo domain.TradingHistoryRepository) DefaultTradingHistoryService {
	return DefaultTradingHistoryService{
		repo: repo,
	}
}

type TradingHistoryService interface {
	// GetUserPositions(ctx context.Context, userID string) ([]*dto.PositionDTO, *errs.AppError)
	GetTradeHistory(ctx context.Context, userID string) ([]*dto.TradeHistoryDTO, *errs.AppError)
}

func (s DefaultTradingHistoryService) GetTradeHistory(ctx context.Context, userID string) ([]*dto.TradeHistoryDTO, *errs.AppError) {
	trades, err := s.repo.GetTradeHistoryByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
//useAccess := middleware.GetUserInfo(ctx)
	//if ok := strings.EqualFold(strings.ToLower(useAccess.Role),"admin"); !ok {
	//	return nil, errs.NewValidationError("userr cannot be empty")
	//}
	
	// Convert domain models to DTOs
	var tradeHistoryDTOs []*dto.TradeHistoryDTO
	for _, trade := range trades {
		tradeHistoryDTOs = append(tradeHistoryDTOs, &dto.TradeHistoryDTO{
			ID:        trade.ID,
			Symbol:    trade.Symbol,
			Volume:    trade.Volume,
			Price:     trade.Price,
			Type:      trade.Type,
			Timestamp: trade.Timestamp,
		})
	}
	return tradeHistoryDTOs, nil
}