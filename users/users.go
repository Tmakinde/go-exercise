package users

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"web-service-gin/database"
)


type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AddUser(ctx *gin.Context) {
	var newUser User
	DB, _ := database.ConnectToDatabase()
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	query := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err := DB.Exec(query, newUser.Username, newUser.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User added successfully"})
}