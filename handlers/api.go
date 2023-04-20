package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
	r := gin.Default()

	posts := r.Group("/posts")
	posts.OPTIONS("", postsOptions)
	posts.GET("", getPosts)
	posts.POST("", addPost)

	post := posts.Group("/:postID")
	post.GET("", getPost)
	post.PATCH("", updatePost)
	// post.DELETE("", deletePost)

	users := r.Group("/users")
	users.OPTIONS("", usersOptions)
	// users.GET("", getUsers)
	// users.POST("",addUser)

	// user := users.Group("/:userID")
	// user.GET("", getUser)
	// user.PATCH("", updateUser)
	// user.DELETE("", deleteUser)

	categories := r.Group("/categories")
	categories.OPTIONS("", categoriesOptions)
	// categories.GET("", getCategories)
	// categories.POST("", addCategory)

	// category := categories.Group("/:categoryID")
	// category.GET("", getCategory)
	// category.PATCH("", updateCategory)
	// category.DELETE("", deleteCategory)

	return r
}
