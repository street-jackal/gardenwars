package models

import "time"

type Plant struct {
	ID              string `json:"id" bson:"ID"`
	Common          string `json:"common,omitempty" bson:"Common,omitempty"`
	Botanical       string `json:"botanical,omitempty" bson:"Botanical,omitempty"`
	Height          []int  `json:"height,omitempty" bson:"Height,omitempty"`
	Characteristics string `json:"characteristics,omitempty" bson:"Characteristics,omitempty"`
	Zones           []int  `json:"zones,omitempty" bson:"Zones,omitempty"`
	Favorited       bool   `json:"favorited,omitempty" bson:"Favorited,omitempty"`
	ImageURL        string `json:"imageURL,omitempty" bson:"ImageURL,omitempty"`

	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
