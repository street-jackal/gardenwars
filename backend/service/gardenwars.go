package gardenwars

import (
	"github.com/gin-gonic/gin"
	"github.com/street-jackal/gardenwars/repository"
)

type Service struct {
	PlantsRepo repository.PlantsRepo
	UsersRepo  repository.UsersRepo
}

type GardenWars interface {
	CreatePlant(c *gin.Context) error
	CreatePlants(c *gin.Context) error

	GetByBotanical(c *gin.Context) error
	GetByCommon(c *gin.Context) error
	GetAllPlants(c *gin.Context) error
	GetAllPlantsForUser(c *gin.Context) error

	CreateUser(c *gin.Context) error
	LoginUser(c *gin.Context) error

	AddUserFavorite(c *gin.Context) error
	RemoveUserFavorite(c *gin.Context) error
}
