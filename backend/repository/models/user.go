package models

import "time"

type User struct {
	ID        string  `json:"id" bson:"ID"`
	Name      string  `json:"name" bson:"Name"`
	Email     string  `json:"email" bson:"Email"`
	Password  string  `json:"password" bson:"Password"`
	Favorites []Plant `json:"favorites,omitempty" bson:"Favorited,omitempty"`

	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
