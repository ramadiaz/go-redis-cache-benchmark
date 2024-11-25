package services

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func CacheOperation(rdb *redis.Client, key string, value string) time.Duration {
	start := time.Now()

	ctx := context.Background()

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Printf("Error setting value: %v\n", err)
	}

	_, err = rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("Error getting value: %v\n", err)
	}

	elapsed := time.Since(start)
	return elapsed
}

func HeatingUpRedis(rdb *redis.Client) {
	ctx := context.Background()

	key := "test"
	value := "test"

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Printf("Error setting value: %v\n", err)
	}

	_, err = rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("Error getting value: %v\n", err)
	}

}
