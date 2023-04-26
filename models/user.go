package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User holds information about users of the system.
type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	FirstName   string             `json:"first_name" bson:"first_name"`
	MiddleName  string             `json:"middle_name" bson:"middle_name"`
	LastName    string             `json:"last_name" bson:"last_name"`
	DateOfBirth time.Time          `json:"date_of_birth" bson:"date_of_birth"`
	UserType    string             `json:"user_type" bson:"user_type"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"-" bson:"password"`
	Bio         string             `json:"bio" bson:"bio"`
	JobTitle    string             `json:"job_title" bson:"job_title"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	Country     string             `json:"country" bson:"country"`
	Sex         string             `json:"sex" bson:"sex"`
}
