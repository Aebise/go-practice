package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"go-practice/db"
	"go-practice/models"
)

// usersOptions is a handler for getting the allowed http methods for the users resource.
func usersOptions(c *gin.Context) {
	methods := []string{http.MethodOptions, http.MethodGet, http.MethodPatch, http.MethodHead, http.MethodPost}
	c.Writer.Header().Set("Allow", strings.Join(methods, " "))
	c.String(http.StatusOK, "")
}

// getUsers is a handler for getting post data.
// It call the GetUsers function from the db package.
// It checks for errors while reading from db.
// It returns 500 if there is error.
// Returns fetched posts if everything goes well.
func getUsers(c *gin.Context) {
	users, err := db.GetUsers()
	if err != nil {
		fmt.Println("error reading users from db: ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	if c.Request.Method == http.MethodHead {
		// if the request method is Head return 200 without body.
		c.String(http.StatusOK, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})

}

// addUser is a handler for adding users.
// It calls the AddUser function from the db package.
// It checks for errors while adding to db.
// It returns 500 if there is an arror.
// Returns 201 if post is added.
func addUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("error binding user data: ", err)
		c.String(http.StatusBadRequest, "")
		return
	}

	user, err := db.AddUser(user)
	if err != nil {
		fmt.Println("error adding user to db: ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	c.String(http.StatusOK, "User created successfully!")
}

// getUser is a handler for getting a single user data using it's id.
// It call the GetUser function from the db package.
// It checks for errors while reading from db.
// It returns 500 if there is error.
// Returns fetched user if everything goes well.
func getUser(c *gin.Context) {
	id := c.Param("userID")

	if id == "" {
		c.String(http.StatusBadRequest, "")
		return
	}

	user, err := db.GetUser(id)
	if err != nil {
		fmt.Println("error reading user from db: ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func updateUser(c *gin.Context) {
	id := c.Param("userID")

	if id == "" {
		c.String(http.StatusBadRequest, "")
		return
	}

	user := models.User{}

	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("error binding user data: ", err)
		c.String(http.StatusBadRequest, "invalid request")
		return
	}

	user, err := db.UpdateUser(user)
	if err != nil {
		fmt.Println("error updating user data on db: ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}
