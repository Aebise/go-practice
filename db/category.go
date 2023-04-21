package db

import (
	"fmt"
	"go-practice/models"
)

func GetCategories() (categories []models.Category, err error) {
	categories = []models.Category{
		{
			ID:          "category1",
			Name:        "Painting",
			Description: "Paint by artists",
		}, {
			ID:          "category2",
			Name:        "Event",
			Description: "Coming events",
		},
	}

	return categories, nil

}

func AddCategory(category models.Category) (models.Category, error) {
	// add to db
	// return the added category including the category ID.

	fmt.Println(category)
	category.ID = "category1"
	return category, nil
}

func GetCategory(id string) (models.Category, error) {
	fmt.Println("id: ", id)
	category := models.Category{
		ID:          "category1",
		Name:        "Painting",
		Description: "Paint by artists",
	}

	return category, nil
}

func UpdateCategory(category models.Category) (models.Category, error) {
	fmt.Println(category)
	// update the category on db
	// return updated data

	return category, nil
}
