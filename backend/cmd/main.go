package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/customsearch/v1"

	"github.com/street-jackal/gardenwars/env"
	"github.com/street-jackal/gardenwars/handlers"
	"github.com/street-jackal/gardenwars/internal/google"
	"github.com/street-jackal/gardenwars/middleware"
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

	// use the default gin router
	router := gin.Default()

	// attach middlwares
	router.Use(middleware.CorsMiddleware())

	// register the routes
	handlers.RegisterPublicRoutes(svc, router)

	router.Run()
}

func initService(ctx context.Context) *gardenwars.Service {
	plantsRepo, err := repository.NewPlantsRepo(ctx)
	if err != nil {
		slog.Error("Failed to initialize the Plants repo", err)
	}

	usersRepo, err := repository.NewUsersRepo(ctx)
	if err != nil {
		slog.Error("Failed to initialize the Users repo", err)
	}

	customsearchService, err := customsearch.NewService(ctx)
	if err != nil {
		slog.Error("Failed to initialize the Custom Search service", err)
	}

	// init the service and return it
	return &gardenwars.Service{
		PlantsRepo:          plantsRepo,
		UsersRepo:           usersRepo,
		CustomSearchService: google.ImageSearch{CustomSearchService: customsearchService},
	}
}
