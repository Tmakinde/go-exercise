package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"fmt"
)

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	fmt.Println("JWT Secret in GenerateToken:", string(jwtSecret))
	claim := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // HS256 is an algorithm that signs the token using hash function and secret key
	return token.SignedString(jwtSecret)
}