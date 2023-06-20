package config

import (
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

var redisClient *redis.Client

func GetRedisClient() *redis.Client {
	databaseIndex, err := strconv.Atoi(os.Getenv("APP_ENV_REDIS_DATABASE"))
	if err != nil {
		panic(err)
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("APP_ENV_REDIS_ADDR"),
		Password: os.Getenv("APP_ENV_REDIS_PASS"),
		DB:       databaseIndex,
	})

	return redisClient
}
