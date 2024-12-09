package cache

import (
	"context"
	"fmt"
	"log"
	"prd_mngt/utils"

	"github.com/go-redis/redis/v8" // Import Redis package
)
var RDB *redis.Client
var ctx = context.Background()

func InitRedis() {
	utils.LoadEnv()
	redisURL := utils.GetEnv("REDIS_URL")

	RDB = redis.NewClient(&redis.Options{
		Addr: redisURL,
	})
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Could not connect to Redis:", err)
	}

	fmt.Println("Connected to Redis")
}

func GetFromCache(key string) (string, error) {
	return RDB.Get(ctx, key).Result()
}

func SetToCache(key string, value string) error {
	return RDB.Set(ctx, key, value, 0).Err()
}
