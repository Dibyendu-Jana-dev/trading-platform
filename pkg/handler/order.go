package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dibyendu/trading_platform/pkg/service"
	"github.com/gorilla/mux"
)
type OrderHandler struct {
	Service service.OrderService
}

// PlaceOrderHandler handles POST requests to place orders.
// CreateUser handles the HTTP POST request to PlaceOrder.
// @Summary Create a new user
// @Description Create a PlaceOrder with the provided details
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} dto.Order "PlaceOrder successfully"
// @Failure 400 {object} string "Invalid request payload"
// @Router /create-user [post]
func (h OrderHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Symbol string  `json:"symbol"`
		Volume float64 `json:"volume"`
		Type   string  `json:"type"`
	}
	ctx := r.Context()

	// Parse the request body into the struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the order data
	if req.Symbol == "" || req.Volume <= 0 || (req.Type != "buy" && req.Type != "sell") {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}

	// Call the service to place the order
	order, appError := h.Service.PlaceOrder(ctx, req.Symbol, req.Volume, req.Type)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
			writeResponse(w, http.StatusOK, order)
		}
}

// @Summary      Delete an order
// @Description  Deletes a limit order specified by its order_id.
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        order_id   path      string  true  "Order ID"
// @Success      200        {object}  map[string]interface{}  "Order deleted successfully"
// @Failure      400        {object}  map[string]string       "Invalid order_id or request"
// @Failure      404        {object}  map[string]string       "Order not found"
// @Failure      500        {object}  map[string]string       "Internal server error"
// @Router       /order/{order_id} [delete]
func (h OrderHandler) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    orderID := vars["order_id"]

    response, appError := h.Service.DeleteOrder(r.Context(), orderID)
    if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
			writeResponse(w, http.StatusOK, response)
		}
}