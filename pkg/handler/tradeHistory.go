package handler

import (
	"net/http"

	"github.com/dibyendu/trading_platform/pkg/service"
	"github.com/gorilla/mux"
)

type TradingHistoryHandler struct {
	Service service.TradingHistoryService
}

// CreateUser handles the HTTP POST request to create a new user.
// @Summary Get a TradeHistory
// @Description Create a new user with the provided details
// @Tags userID
// @Accept json
// @Produce json
// @Failure 400 {object} string "Invalid request payload"
// @Router /create-user [post]
func (h TradingHistoryHandler) GetTradeHistory(w http.ResponseWriter, r *http.Request) {
	// Assume userID is retrieved from the JWT token (or as a placeholder here).
	userID := mux.Vars(r)["user_id"]
	ctx     := r.Context()

	tradeHistory, appError := h.Service.GetTradeHistory(ctx, userID)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, tradeHistory)
	}
}