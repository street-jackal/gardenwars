package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/street-jackal/gardenwars/env"
	"github.com/street-jackal/gardenwars/handlers"
	"github.com/street-jackal/gardenwars/repository"
	gardenwars "github.com/street-jackal/gardenwars/service"
)

func main() {
	// make the init logic time-bound
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// init any env vars
	env.Init()

	// init the service
	svc := initService(ctx)

	router := gin.Default()
	handlers.RegisterPublicRoutes(svc, router)

	router.Run()
}

func initService(ctx context.Context) *gardenwars.Service {
	plantsRepo, err := repository.NewPlantsRepo(ctx)
	if err != nil {
		slog.Error("Failed to initialize the Plants repo", err)
	}

	// init the service and return it
	return &gardenwars.Service{
		PlantsRepo: plantsRepo,
	}
}
