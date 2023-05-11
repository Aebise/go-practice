package handlers

import (
	"go-practice/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	api.POST("/login", loginHandler)
	api.POST("/users", addUser)

	protected := api.Group("/protected")
	protected.Use(middlewares.JwtAuthMiddleware())

	posts := protected.Group("/posts")
	posts.OPTIONS("", postsOptions)
	posts.GET("", getPosts)
	posts.POST("", addPost)

	post := posts.Group("/:postID")
	post.GET("", getPost)
	post.PATCH("", updatePost)
	// post.DELETE("", deletePost)

	users := protected.Group("/users")
	users.OPTIONS("", usersOptions)
	users.GET("", getUsers)

	user := users.Group("/:userID")
	user.GET("", getUser)
	user.PATCH("", updateUser)
	// user.DELETE("", deleteUser)

	categories := protected.Group("/categories")
	categories.OPTIONS("", categoriesOptions)
	categories.GET("", getCategories)
	categories.POST("", addCategory)

	category := categories.Group("/:categoryID")
	category.GET("", getCategory)
	category.PATCH("", updateCategory)
	// category.DELETE("", deleteCategory)

	return r
}
