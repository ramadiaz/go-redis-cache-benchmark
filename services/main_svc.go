package services

import (
	"fmt"

	"github.com/cheggaaa/pb/v3"
	"github.com/go-redis/redis/v8"
)

func RunTest(rdb *redis.Client, testName string, iterations int) ([]string, []float64) {
	var keys []string
	var times []float64

	bar := pb.StartNew(iterations)

	for i := 1; i <= iterations; i++ {
		key := fmt.Sprintf("%s_key_%d", testName, i)
		value := fmt.Sprintf("%s_value_%d", testName, i)
		responseTime := CacheOperation(rdb, key, value)

		keys = append(keys, key)
		times = append(times, responseTime.Seconds()*1000) 

		bar.Increment()
	}

	bar.Finish()

	return keys, times
}