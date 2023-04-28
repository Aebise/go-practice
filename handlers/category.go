package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-practice/db"
	"go-practice/models"
)

// categoriesOptions is a handler for getting the allowed http methods for the categories resource.
func categoriesOptions(c *gin.Context) {
	methods := []string{http.MethodOptions, http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPatch}
	c.Writer.Header().Set("Allow", strings.Join(methods, " "))
	c.String(http.StatusOK, "")
}

// getCategories is a handler for getting categories data.
// It call the GetCategories function from the db package.
// It checks for errors while reading from db.
// It returns 500 if there is error.
// Returns fetched categories if everything goes well.
func getCategories(c *gin.Context) {

	categories, err := db.GetCategories()
	if err != nil {
		fmt.Println("error reading categories from db: ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	if c.Request.Method == http.MethodHead {
		// if the request method is Head return 200 without body.
		c.String(http.StatusOK, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})

}

// addCategory is a handler for adding category.
// It calls the AddCategory function from the db package.
// It checks for errors while adding to db.
// It returns 500 if there is an arror.
// Returns 201 if category is added.
func addCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBind(&category); err != nil {
		fmt.Println("error binding category data: ", err)
		c.String(http.StatusBadRequest, "invalid request")
		return
	}

	category, err := db.AddCategory(category)
	if err != nil {
		fmt.Println("error creating category: ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	fmt.Println(category)
	c.String(http.StatusCreated, "Category created successfully!")
}

// getCategory is a handler for getting a single category data using it's id.
// It call the GetCategory function from the db package.
// It checks for errors while reading from db.
// It returns 500 if there is error.
// Returns fetched category if everything goes well.
func getCategory(c *gin.Context) {
	id := c.Param("categoryID")

	if id == "" {
		c.String(http.StatusBadRequest, "")
		return
	}

	category, err := db.GetCategory(id)
	if err != nil {
		fmt.Println("error reading category from db: ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

// updateCategory is a handler for updating a single category data using it's id.
// It call the UpdateCategory function from the db package.
// It checks for errors while updating on db.
// It returns 500 if there is error.
// Returns updated category if everything goes well.
func updateCategory(c *gin.Context) {
	id := c.Param("categoryID")

	if id == "" {
		c.String(http.StatusBadRequest, "")
		return
	}

	category := models.Category{}
	if err := c.ShouldBind(&category); err != nil {
		fmt.Println("error binding category data: ", err)
		c.String(http.StatusBadRequest, "invalid request")
		return
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("error parsing id: ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	category.ID = objId

	category, err = db.UpdateCategory(category)
	if err != nil {
		fmt.Println("error updating category on db: ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}
