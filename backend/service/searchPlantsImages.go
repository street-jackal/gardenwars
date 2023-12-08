package gardenwars

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func (svc *Service) SearchPlantsImages(c *gin.Context) {
	type SearchPlantsImagesRequest struct {
		PlantIDs []string `json:"plantIDs"`
	}

	req := &SearchPlantsImagesRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plants, err := svc.PlantsRepo.GetMultiple(context.Background(), req.PlantIDs)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			slog.Warn("No Plants found for search", err)

			c.JSONP(http.StatusOK, nil)

			return
		}
	}

	var updated int
	for _, plant := range plants {
		imageURLs, err := svc.CustomSearchService.Search(plant.Botanical)
		if err != nil {
			slog.Error("Failed to search for Plant image", err)
		}

		c.JSONP(
			http.StatusUnauthorized, gin.H{"error": err.Error()},
		)

		if len(imageURLs) > 0 {
			if err := svc.PlantsRepo.SetImage(c, plant.ID, imageURLs[0]); err != nil {
				slog.Error("Failed to set Plant image", err)
			}

			updated++
		}
	}

	c.JSONP(
		http.StatusOK,
		gin.H{
			"message": fmt.Sprintf("Successfully searched for %d Plant images", updated),
		},
	)
}
