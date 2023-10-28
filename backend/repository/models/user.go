package models

import "time"

type User struct {
	ID        string   `json:"id" bson:"ID"`
	Email     string   `json:"email" bson:"Email"`
	Password  string   `json:"password" bson:"Password"`
	Favorites []string `json:"favorites,omitempty" bson:"Favorites,omitempty"`

	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
