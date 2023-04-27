package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Post is a model that holds post data.
type Post struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	ImageURL    string             `json:"image_url" bson:"image_url"`
	VideoURL    string             `json:"video_url" bson:"video_url"`
	Views       int64              `json:"views" bson:"views"`
	Like        int64              `json:"like" bson:"like"`
	PosterID    string             `json:"poster_id" bson:"poster_id"`
	CategoryID  string             `json:"category_id" bson:"category_id"`
	Status      string             `json:"status" bson:"status"`
	Description string             `json:"description" bson:"description"`
}
