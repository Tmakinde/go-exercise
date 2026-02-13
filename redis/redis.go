package redis

import (
	"github.com/redis/go-redis/v9"
	"github.com/joho/godotenv"
	"fmt"
	"os"
	"strconv"
	"context"
)

var ctx = context.Background()

func ConnectToRedis() (*redis.Client, error) {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, err
	}
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		fmt.Println("Error converting REDIS_DB to integer:", err)
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
        Addr:    fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
        Password: os.Getenv("REDIS_PASSWORD"),
        DB:       db,
    })

	return rdb, nil
}

func SetValue(rdb *redis.Client, key string, value string) error {
	err := rdb.Set(ctx, key, value, 0).Err()
	return err
}

func GetValue(rdb *redis.Client, key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	return val, err
}