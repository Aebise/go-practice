package models

import (
	"time"
)

// User holds information about users of the system.
type User struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	MiddleName  string    `json:"middle_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	UserType    string    `json:"user_type"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Bio         string    `json:"bio"`
	JobTitle    string    `json:"job_title"`
	PhoneNumber string    `json:"phone_number"`
	Country     string    `json:"country"`
	Sex         string    `json:"sex"`
}
