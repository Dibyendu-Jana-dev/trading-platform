package handler

import (
	"net/http"
	"github.com/dibyendu/trading_platform/pkg/service"
	"github.com/gorilla/mux"
)

type MarketDataHandler struct {
	Service service.MarketDataService
}

// GetUser handles the HTTP GET request to fetch user details.
// @Security ApiKeyAuth
// @Summary Fetch user details
// @Description Fetch GetMarketData
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "symbol"
// @Security bearerToken
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unauthorized access"
// @Failure 404 {object} string "User not found"
// @Router /user/get-user [get]
func (m MarketDataHandler) GetMarketData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    symbol := vars["symbol"]
	ctx     := r.Context()

    marketData, appError := m.Service.GetMarketData(ctx, symbol)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, marketData)
	}
}