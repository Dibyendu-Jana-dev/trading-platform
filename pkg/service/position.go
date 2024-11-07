package service

import (
	"context"

	"github.com/dibyendu/trading_platform/lib/errs"
	"github.com/dibyendu/trading_platform/pkg/domain"
	"github.com/dibyendu/trading_platform/pkg/dto"
)

type DefaultPositionService struct {
	repo domain.PositionRepository
}

func NewPositionService(repo domain.PositionRepository) DefaultPositionService {
	return DefaultPositionService{
		repo: repo,
	}
}

type PositionService interface {
	GetUserPositions(ctx context.Context, userID string) ([]*dto.PositionDTO, *errs.AppError)
}


func (s DefaultPositionService) GetUserPositions(ctx context.Context, userID string) ([]*dto.PositionDTO, *errs.AppError) {
	positions, err := s.repo.GetUserPositions(ctx, userID)
	if err != nil {
		return nil, err
	}

	//useAccess := middleware.GetUserInfo(ctx)
	//if ok := strings.EqualFold(strings.ToLower(useAccess.Role),"admin"); !ok {
	//	return nil, errs.NewValidationError("userr cannot be empty")
	//}

	var positionsDTO []*dto.PositionDTO
	for _, position := range positions {
		positionsDTO = append(positionsDTO, &dto.PositionDTO{
			Symbol:    position.Symbol,
			Quantity:  position.Quantity,
			AvgPrice:  position.AvgPrice,
			Timestamp: position.Timestamp,
		})
	}

	return positionsDTO, nil
}