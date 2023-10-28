package gardenwars

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/street-jackal/gardenwars/repository/models"
	"github.com/street-jackal/gardenwars/service/types/responses"
	"go.mongodb.org/mongo-driver/mongo"
)

func (svc *Service) GetAllPlantsForUser(c *gin.Context) {
	type GetAllPlantsForUserRequest struct {
		UserID string `json:"userID"`
	}

	req := &GetAllPlantsForUserRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := svc.UsersRepo.Get(context.Background(), req.UserID)
	if err != nil {
		slog.Error("Failed to retrieve User", err)

		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, responses.BaseResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "User does not exist",
			})

			return
		}
	}

	plants, err := svc.PlantsRepo.GetMultiple(context.Background(), user.Favorites)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			slog.Warn("No Plants found for User", err)

			c.JSONP(http.StatusOK, []models.Plant{})

			return
		}

	}

	c.JSONP(http.StatusOK, plants)
}
