package main

import (
	"fmt"
	"go-redis-cache-benchmark/config"
	"go-redis-cache-benchmark/services"
	"time"
)

func main() {

	rdb := config.ConnectRedis()
	defer rdb.Close()

	services.HeatingUpRedis(rdb)

	testName := "redis_test"
	iterations := 10000

	fmt.Println("Running Test 1...")
	keys1, times1 := services.RunTest(rdb, testName+"_1", iterations)
	services.SaveResultsToCSV("test1_results.csv", keys1, times1)
	fmt.Println("Test 1 completed, results saved to test1_results.csv")

	time.Sleep(5 * time.Second)

	fmt.Println("Running Test 2...")
	keys2, times2 := services.RunTest(rdb, testName+"_2", iterations)
	services.SaveResultsToCSV("test2_results.csv", keys2, times2)
	fmt.Println("Test 2 completed, results saved to test2_results.csv")

	fmt.Println("Saving final result to final_result.csv")

	err := services.SaveFinalResultToCSV("final_result.csv", times1, times2)
	if err != nil {
		fmt.Printf("Error saving final_result.csv: %s", err.Error())
	}

	fmt.Println("Benchmark successfully executed")
}
