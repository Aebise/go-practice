package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Category is a model that holds data about categories of events.
type Category struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
}
