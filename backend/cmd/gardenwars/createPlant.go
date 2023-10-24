package gardenwars

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/street-jackal/gardenwars/repository/models"
)

func (svc *Service) CreatePlant(c *gin.Context) {
	var plant models.Plant

	if err := c.ShouldBindJSON(&plant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	if err := svc.PlantsRepo.Insert(c.Request.Context(), &plant); err != nil {
		slog.Error("Failed to insert a Plant", err)
	}
	
	c.JSONP(http.StatusOK, plant)
}
