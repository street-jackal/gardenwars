package gardenwars

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/street-jackal/gardenwars/service/types/responses"
)

func (svc *Service) RemoveUserFavorite(c *gin.Context) {
	type RemoveUserFavoriteRequest struct {
		UserID  string `json:"userID"`
		PlantID string `json:"plantID"`
	}

	req := &RemoveUserFavoriteRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Failed to bind the request body to a struct", err)

		c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		})

		return
	}

	if err := svc.UsersRepo.RemoveFavorite(
		c.Request.Context(),
		req.UserID,
		req.PlantID,
	); err != nil {
		slog.Error("Failed to remove a favorite Plant", err)
	}

	c.JSONP(http.StatusOK, responses.BaseResponse{
		Success: true,
		Message: "Successfully removed a favorite Plant",
	})
}
