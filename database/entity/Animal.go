package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type AnimalEntity struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Species string             `json:"species" bson:"species,omitempty"`
}
