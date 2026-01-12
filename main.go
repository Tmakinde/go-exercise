package main

import (
	"web-service-gin/exercises"
	"web-service-gin/users"
	"github.com/gin-gonic/gin"
	"web-service-gin/middleware"
	"os"
)

func main() {
	router := gin.Default()
    router.GET("/albums", exercises.GetAlbums)
	router.POST("/albums", exercises.PostAlbum)
	router.POST("/users/register", users.AddUser)
	router.POST("/users/login", users.Login)
	protected := router.Group("/users")

	protected.Use(middleware.AuthMiddleware([]byte(os.Getenv("JWT_SECRET"))))
	{
		protected.GET("/profile", users.GetProfile)
		protected.PUT("/reset-password", users.UpdatePassword)
	}

	// Run the server
    router.Run("localhost:8080")
}
