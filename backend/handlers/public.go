package handlers

import (
	"github.com/gin-gonic/gin"
	gardenwars "github.com/street-jackal/gardenwars/service"
)

func RegisterPublicRoutes(svc *gardenwars.Service, r *gin.Engine) {
	// plants endpoints
	r.GET("/plants/botanical/get", svc.GetByBotanical)
	r.GET("/plants/getAll", svc.GetAllPlants)

	r.POST("/plants/create", svc.CreatePlant)
	r.POST("/plants/createMany", svc.CreatePlants)
}
