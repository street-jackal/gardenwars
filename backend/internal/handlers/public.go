package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/street-jackal/gardenwars/cmd/gardenwars"
)

func RegisterPublicRoutes(svc *gardenwars.Service, r *gin.Engine) {
	// plants endpoints
	r.GET("/plants/botanical/get", svc.GetByBotanical)

	r.POST("/plants/create", svc.CreatePlant)
	r.POST("/plants/createMany", svc.CreatePlants)
}
