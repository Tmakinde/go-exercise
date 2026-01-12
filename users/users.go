package users

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"web-service-gin/database"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	"web-service-gin/utils"
	"fmt"
)

var DB, _ = database.ConnectToDatabase()

type User struct {
	ID	   int    `json:"id"`
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
	hashedPassword, _ := hashPassword(newUser.Password)
	_, err := DB.Exec(query, newUser.Username, hashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User added successfully"})
}

func Login(ctx *gin.Context) {
	var loginUser User
	if err := ctx.BindJSON(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	query := `SELECT password FROM users where username=$1`

	result := DB.QueryRow(query, loginUser.Username)

	// Store hashed password from DB here
	var hashedPassword string

	err := result.Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(loginUser.Password),
	)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	token, _ := utils.GenerateToken(loginUser.Username)
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func GetProfile(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	ctx.JSON(http.StatusOK, gin.H{"username": username})
}

func UpdatePassword(ctx *gin.Context) {
	type pwdUpdate struct {
		Password string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	var pwd pwdUpdate
	if err := ctx.BindJSON(&pwd); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Println("Password:", pwd.Password)
	fmt.Println("Confirm Password:", pwd.ConfirmPassword)
	if pwd.Password != pwd.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	hashPassword, err := hashPassword(pwd.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	username, _ := ctx.Get("username")
	
	query := `UPDATE users SET password=$1 WHERE username=$2`
	_, err = DB.Exec(query, hashPassword, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

