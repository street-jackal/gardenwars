package handlers

import (
	"github.com/gin-gonic/gin"
	gardenwars "github.com/street-jackal/gardenwars/service"
)

func RegisterPublicRoutes(svc *gardenwars.Service, r *gin.Engine) {
	// plants endpoints
	r.GET("/plants/botanical/get", svc.GetByBotanical)
	r.GET("/plants/getAll", svc.GetAllPlants)
	r.POST("/plants/getAllForUser", svc.GetAllPlantsForUser)

	r.POST("/plants/create", svc.CreatePlant)
	r.POST("/plants/createMany", svc.CreatePlants)

	r.POST("/plants/images/update", svc.SearchPlantsImages)

	// users endpoints
	r.POST("/users/signup", svc.CreateUser)
	r.POST("/users/login", svc.LoginUser)
	r.POST("/users/favorites/add", svc.AddUserFavorite)
	r.POST("/users/favorites/remove", svc.RemoveUserFavorite)
}
