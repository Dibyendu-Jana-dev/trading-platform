package service

import (
	"context"

	"github.com/dibyendu/trading_platform/lib/errs"

	"github.com/dibyendu/trading_platform/pkg/domain"
	"github.com/dibyendu/trading_platform/pkg/dto"
)


type DefaultMarketDataService struct {
	repo domain.MarketDataRepository
}

func NewMarketDataService(repo domain.MarketDataRepository) DefaultMarketDataService {
	return DefaultMarketDataService{
		repo: repo,
	}
}

type MarketDataService interface {
	// CreateUser(ctx context.Context, request dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError)
	// SignIn(ctx context.Context, request dto.CreateUserRequest) (*dto.UserLoginResponse, *errs.AppError)
	// GetUser(ctx context.Context, request dto.GetUserRequest) (*dto.GetUserResponse, *errs.AppError)
	// PlaceOrder(ctx context.Context, request dto.Order) (*dto.Order, *errs.AppError)
	GetMarketData(ctx context.Context, symbol string) (*dto.MarketDataResponse, *errs.AppError)
}


func (s DefaultMarketDataService) GetMarketData(ctx context.Context, symbol string) (*dto.MarketDataResponse, *errs.AppError) {
	var (
		response    *dto.MarketDataResponse
	)
	//useAccess := middleware.GetUserInfo(ctx)
	//if ok := strings.EqualFold(strings.ToLower(useAccess.Role),"admin"); !ok {
	//	return nil, errs.NewValidationError("userr cannot be empty")
	//}
	
	data, err := s.repo.GetMarketData(ctx, symbol)

	if err != nil {
		return nil, errs.NewNoContentError("u")
	}
	response = data.ToDto()
	return response, nil
}