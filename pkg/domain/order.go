package domain

import (
	"context"
	"time"

	"github.com/dibyendu/trading_platform/lib/errs"
	"github.com/dibyendu/trading_platform/pkg/dto"
)

type OrderRepository interface {
	PlaceOrder(ctx context.Context, req Order) (*Order, *errs.AppError)
	DeleteOrder(ctx context.Context, orderId string)(*DeleteOrderResponse, *errs.AppError)
}

type Order struct {
	ID        string    `json:"id" bson :"id"`
	Symbol    string    `json:"symbol" bson:"symbol"`
	Volume    float64   `json:"volume" bson:"volume"`
	Type      string    `json:"type" bson:"type"`
	Price     float64   `json:"price" bson:"price"`
	Status    string    `json:"status" bson:"status"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}


func (r Order) ToDto() *dto.Order{
	return &dto.Order{
		ID: r.ID,
		Symbol: r.Symbol,
		Volume: r.Volume,
		Type: r.Type,
		Price: r.Price,
		Status: r.Status,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}


type DeleteOrderResponse struct {
    Message string `json:"message"`
    OrderID string `json:"order_id"`
}

func (r DeleteOrderResponse) ToDto() *dto.DeleteOrderResponse{
	return &dto.DeleteOrderResponse{
		Message: r.Message,
		OrderID: r.OrderID,
	}
}