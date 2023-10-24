package gardenwars

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (svc *Service) GetByBotanical(c *gin.Context) {
	type GetByBotanicalReq struct {
		Botanical string `json:"botanical"`
	}

	var req GetByBotanicalReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plants, err := svc.PlantsRepo.GetByBotanical(context.Background(), req.Botanical)
	if err != nil {
		slog.Error("Failed to retrieve a Plant", err)
	}

	c.JSONP(http.StatusOK, plants)
}
