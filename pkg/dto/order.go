package dto
import "time"

type Order struct {
	ID        string    `json:"id"`
	Symbol    string    `json:"symbol"`
	Volume    float64   `json:"volume"`
	Type      string    `json:"type"`
	Price     float64   `json:"price"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteOrderResponse struct {
    Message string `json:"message"`
    OrderID string `json:"order_id"`
}
