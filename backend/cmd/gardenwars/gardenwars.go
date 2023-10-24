package gardenwars

import (
	"github.com/gin-gonic/gin"
	"github.com/street-jackal/gardenwars/repository"
)

type Service struct {
	PlantsRepo repository.PlantsRepo
}

type GardenWars interface {
	CreatePlant(c *gin.Context) error
	CreatePlants(c *gin.Context) error

	GetByBotanical(c *gin.Context) error
	GetByCommon(c *gin.Context) error
}
