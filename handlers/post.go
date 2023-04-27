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

// postsOptions is a handler for getting the allowed http methods for the posts resource.
func postsOptions(c *gin.Context) {
	methods := []string{http.MethodOptions, http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPatch}
	c.Writer.Header().Set("Allow", strings.Join(methods, " "))
	c.String(http.StatusOK, "thank you")
}

// getPosts is a handler for getting post data.
// It call the GetPosts function from the db package.
// It checks for errors while reading from db.
// It returns 500 if there is error.
// Returns fetched posts if everything goes well.
func getPosts(c *gin.Context) {

	posts, err := db.GetPosts()
	if err != nil {
		fmt.Println("error reading posts from db: ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	if c.Request.Method == http.MethodHead {
		// if the request method is Head return 200 without body.
		c.String(http.StatusOK, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}

// addPost is a handler for adding posts.
// It calls the AddPost function from the db package.
// It checks for errors while adding to db.
// It returns 500 if there is an arror.
// Returns 201 if post is added.
func addPost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBind(&post); err != nil {
		fmt.Println("error binding post data: ", err)
		c.String(http.StatusBadRequest, "invalid request")
		return
	}

	post, err := db.AddPost(post)
	if err != nil {
		fmt.Println("error creating post: ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	fmt.Println(post)
	c.String(http.StatusCreated, "Post created successfully!")
}

// getPost is a handler for getting a single post data using it's id.
// It call the GetPost function from the db package.
// It checks for errors while reading from db.
// It returns 500 if there is error.
// Returns fetched post if everything goes well.
func getPost(c *gin.Context) {
	id := c.Param("postID")

	if id == "" {
		c.String(http.StatusBadRequest, "")
		return
	}

	post, err := db.GetPost(id)
	if err != nil {
		fmt.Println("error reading post from db: ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// updatePost is a handler for updating a single post data using it's id.
// It call the UpdatePost function from the db package.
// It checks for errors while updating on db.
// It returns 500 if there is error.
// Returns updated post if everything goes well.
func updatePost(c *gin.Context) {
	id := c.Param("postID")

	if id == "" {
		c.String(http.StatusBadRequest, "")
		return
	}

	post := models.Post{}
	if err := c.ShouldBind(&post); err != nil {
		fmt.Println("error binding post data: ", err)
		c.String(http.StatusBadRequest, "invalid request")
		return
	}
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("error parsing ID; ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}
	post.ID = objId

	post, err = db.UpdatePost(post)
	if err != nil {
		fmt.Println("error updating post on db: ", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}
