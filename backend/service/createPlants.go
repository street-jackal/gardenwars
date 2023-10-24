package gardenwars

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/street-jackal/gardenwars/repository/models"
)

func (svc *Service) CreatePlants(c *gin.Context) {
	var plants = make([]*models.Plant, 0)

	if err := c.ShouldBindJSON(&plants); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := svc.PlantsRepo.InsertMany(c.Request.Context(), plants); err != nil {
		slog.Error("Failed to insert multiple Plants", err)
	}

	c.JSONP(http.StatusOK, "Plants created")
}
