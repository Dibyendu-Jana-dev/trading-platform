package service

import (
	"context"
	"fmt"
	"time"

	"github.com/dibyendu/trading_platform/lib/errs"
	"github.com/dibyendu/trading_platform/pkg/domain"
	"github.com/dibyendu/trading_platform/pkg/dto"
	fetcherservice "github.com/dibyendu/trading_platform/pkg/httpClient/fetcherService"
)

type DefaultOrderService struct {
	repo domain.OrderRepository
}

func NewOrderService(repo domain.OrderRepository) DefaultOrderService {
	return DefaultOrderService{
		repo: repo,
	}
}

type OrderService interface {
	PlaceOrder(ctx context.Context, symbol string, volume float64, orderType string) (*dto.Order, *errs.AppError)
	DeleteOrder(ctx context.Context, orderId string)(*dto.DeleteOrderResponse, *errs.AppError)
}

func (s DefaultOrderService) PlaceOrder(ctx context.Context, symbol string, volume float64, orderType string) (*dto.Order, *errs.AppError) {
	// Fetch the latest market data for the given symbol
	marketData, err := fetcherservice.GetMarketDataGetMarketData(symbol)
	if err != nil {
		return &dto.Order{}, errs.NewNoContentError("failed to fetch market data: %v")
	}

	//useAccess := middleware.GetUserInfo(ctx)
	//if ok := strings.EqualFold(strings.ToLower(useAccess.Role),"admin"); !ok {
	//	return nil, errs.NewValidationError("userr cannot be empty")
	//}

	// Calculate the price (based on the fetched market data)
	var price float64
	if orderType == "buy" {
		price = marketData.AskPrice
	} else if orderType == "sell" {
		price = marketData.BidPrice
	} else {
		return &dto.Order{}, errs.NewNotFoundError("invalid order type")
	}

	// Create the order object
	order := domain.Order{
		ID:        fmt.Sprintf("order-%d", time.Now().UnixNano()), // Unique ID based on timestamp
		Symbol:    symbol,
		Volume:    volume,
		Type:      orderType,
		Price:     price,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save the order to the database
	savedOrder, ers := s.repo.PlaceOrder(ctx, order)
	if ers != nil {
		return &dto.Order{}, errs.NewNoContentError("failed to save order")
	}
	response := savedOrder.ToDto()
	return response, nil
}


func (s DefaultOrderService) DeleteOrder(ctx context.Context, orderID string) (*dto.DeleteOrderResponse, *errs.AppError) {
    var (
		res *dto.DeleteOrderResponse
	)

	//useAccess := middleware.GetUserInfo(ctx)
	//if ok := strings.EqualFold(strings.ToLower(useAccess.Role),"admin"); !ok {
	//	return nil, errs.NewValidationError("userr cannot be empty")
	//}
	response, err := s.repo.DeleteOrder(ctx, orderID)
    if err != nil {
        return &dto.DeleteOrderResponse{}, errs.NewNoContentError("failed to Delete order")
     }
    res = response.ToDto()
	return res, nil
}