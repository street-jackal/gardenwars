package main

import (
	"encoding/json"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/street-jackal/gardenwars/repository/models"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func main() {
	// Open the input file
	inputFile, err := os.Open("plants.json")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// Read the input file
	inputBytes, err := io.ReadAll(inputFile)
	if err != nil {
		panic(err)
	}

	// Unmarshal the input JSON into an array of Plant structs
	var plants []Plant2
	err = json.Unmarshal(inputBytes, &plants)
	if err != nil {
		panic(err)
	}

	newPlants := make([]models.Plant, 0, len(plants))
	for _, plant := range plants {

		height := make([]int, 0, len(plant.Height))
		zones := make([]int, 0, len(plant.Zones))

		for _, h := range plant.Height {
			h = clearString(h)
			hstring, err := strconv.Atoi(h)
			if err != nil {
				panic(err)
			}
			height = append(height, hstring)
		}

		for _, z := range plant.Zones {
			z = clearString(z)
			zstring, err := strconv.Atoi(z)
			if err != nil {
				panic(err)
			}
			zones = append(zones, zstring)
		}

		newPlants = append(newPlants, models.Plant{
			ID:              plant.ID,
			Common:          plant.Common,
			Botanical:       plant.Botanical,
			Height:          height,
			Characteristics: plant.Characteristics,
			Zones:           zones,
			Favorited:       plant.Favorited,
		})
	}

	// Marshal the modified plants array into JSON
	outputBytes, err := json.Marshal(newPlants)
	if err != nil {
		panic(err)
	}

	// Write the output JSON to a new file
	err = os.WriteFile("formatted_plants.json", outputBytes, 0644)
	if err != nil {
		panic(err)
	}
}

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

type Plant2 struct {
	ID              string   `json:"id" bson:"ID"`
	Common          string   `json:"common,omitempty" bson:"Common,omitempty"`
	Botanical       string   `json:"botanical,omitempty" bson:"Botanical,omitempty"`
	Height          []string `json:"height,omitempty" bson:"Height,omitempty"`
	Characteristics string   `json:"characteristics,omitempty" bson:"Characteristics,omitempty"`
	Zones           []string `json:"zones,omitempty" bson:"Zones,omitempty"`
	Favorited       bool     `json:"favorited,omitempty" bson:"Favorited,omitempty"`

	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
