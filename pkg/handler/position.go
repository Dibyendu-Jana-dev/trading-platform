package handler

import (
	"net/http"
	"strings"

	"github.com/dibyendu/trading_platform/lib/errs"
	"github.com/dibyendu/trading_platform/pkg/middleware"
	"github.com/dibyendu/trading_platform/pkg/service"
)

type PositionHandler struct {
	Service service.PositionService
}

// GetUser handles the HTTP GET request to fetch user details.
// @Security ApiKeyAuth
// @Summary Fetch GetPositions
// @Description Fetch details of a user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Security bearerToken
// @Success 200 {object} dto.PositionDTO "GetPositions fetched successfully"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unauthorized access"
// @Failure 404 {object} string "User not found"
// @Router /user/get-user [get]

func (h PositionHandler) GetPositions(w http.ResponseWriter, r *http.Request) {
	// Validate JWT and get user ID from token
	ctx := r.Context()
	userInfo := middleware.GetUserInfo(ctx)
	if strings.EqualFold(strings.ToLower(userInfo.AuthToken), "invalid") {
		writeResponse(w, http.StatusUnauthorized, errs.NewValidationError("unauthorized Access for get user").AsMessage())
		return
	}

	// Fetch positions using the service
	positions, err := h.Service.GetUserPositions(ctx, userInfo.UserId)
	if err != nil {
		if err.Code == http.StatusNoContent {
			writeResponseNoContent(w, err.Code)
		} else {
			writeResponse(w, err.Code, err.AsMessage())
		}
	} else {
		writeResponse(w, http.StatusOK, positions)
	}
}