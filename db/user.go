package db

import (
	"fmt"
	"go-practice/models"
)

func GetUsers() ([]models.User, error) {

	users := []models.User{
		{
			ID:          "User1",
			FirstName:   "User",
			MiddleName:  "Father",
			LastName:    "LastName",
			UserType:    "Admin",
			Email:       "user1@gmail.com",
			Bio:         "IT",
			JobTitle:    "Technician",
			PhoneNumber: "+2519-45-56-43",
			Country:     "Ethiopia",
			Sex:         "Female",
		},
		{
			ID:          "User2",
			FirstName:   "Usertwo",
			MiddleName:  "Father",
			LastName:    "LastName",
			UserType:    "Poster",
			Email:       "user2@gmail.com",
			Bio:         "Artist",
			JobTitle:    "Artist",
			PhoneNumber: "+2519-67-98-14",
			Country:     "Ethiopia",
			Sex:         "Male",
		},
	}

	return users, nil
}

func AddUser(user models.User) (models.User, error) {
	fmt.Println(user)
	user.ID = "user1"

	return user, nil
}

func GetUser(id string) (models.User, error) {
	fmt.Println("id : ", id)

	user := models.User{
		ID:          "User1",
		FirstName:   "User",
		MiddleName:  "Father",
		LastName:    "LastName",
		UserType:    "Admin",
		Email:       "user1@gmail.com",
		Bio:         "IT",
		JobTitle:    "Technician",
		PhoneNumber: "+2519-45-56-43",
		Country:     "Ethiopia",
		Sex:         "Female",
	}

	return user, nil

}

func UpdateUser(user models.User) (models.User, error) {
	fmt.Println(user)
	// update the user on db
	// return updated data

	return user, nil
}
