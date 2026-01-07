package database

import (
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/joho/godotenv"
	"fmt"
	"os"
)

func ConnectToDatabase() (*sql.DB, error) {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, err
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, err
	}

	fmt.Println("Successfully connected to database")
	return db, nil
}