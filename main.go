package main

import (
	//"first/controllers"
	"first/controllers"
	"first/initializers"

	"github.com/gin-gonic/gin"
)
func init() {
	initializers.LoadEnvVariables()
}
func main() {
	
	// Create a Gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

 	r.POST("/post", controllers.CreatePostHandler)
	r.GET("/all-post", controllers.GetAllPostsHandler)
	r.GET("/post/:id", controllers.GetPostByIdHandler)
	r.DELETE("/post/:id", controllers.DeletePostByIdHandler)

	r.Run()
}