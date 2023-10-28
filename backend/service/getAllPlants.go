package gardenwars

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (svc *Service) GetAllPlants(c *gin.Context) {

	plants, err := svc.PlantsRepo.GetAll(context.Background())
	if err != nil {
		slog.Error("Failed to retrieve Plants", err)
		return
	}

	c.JSONP(http.StatusOK, plants)
}
