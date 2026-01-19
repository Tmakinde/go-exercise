package main

import (
	"web-service-gin/exercises"
	"web-service-gin/users"
	"github.com/gin-gonic/gin"
	"web-service-gin/middleware"
	"web-service-gin/filemanager"
	"os"
	"web-service-gin/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	metrics.Init()
	router := gin.Default()

	// Prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Public routes
    router.GET("/albums", exercises.GetAlbums)
	router.POST("/albums", exercises.PostAlbum)
	router.POST("/users/register", users.AddUser)
	router.POST("/users/login", users.Login)
	protected := router.Group("/users")

	protected.Use(middleware.AuthMiddleware([]byte(os.Getenv("JWT_SECRET"))))
	{
		protected.GET("/profile", users.GetProfile)
		protected.PUT("/reset-password", users.UpdatePassword)
		protected.POST("/upload", filemanager.UploadFileToS3)
	}

	// Run the server
    router.Run("localhost:8080")
}
