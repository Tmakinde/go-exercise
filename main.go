package main

import (
	"web-service-gin/exercises"
	"web-service-gin/users"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
    router.GET("/albums", exercises.GetAlbums)
	router.POST("/albums", exercises.PostAlbum)
	router.POST("/users", users.AddUser)
	// Run the server
    router.Run("localhost:8080")
}
