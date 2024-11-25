package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

func PearsonCorrelation(x, y []float64) float64 {
	n := len(x)
	if n != len(y) || n == 0 {
		panic("Datasets must have the same non-zero length.")
	}

	var sumX, sumY, sumXY, sumX2, sumY2 float64
	for i := 0; i < n; i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
		sumY2 += y[i] * y[i]
	}

	numerator := float64(n)*sumXY - sumX*sumY
	denominator := math.Sqrt((float64(n)*sumX2 - sumX*sumX) * (float64(n)*sumY2 - sumY*sumY))
	if denominator == 0 {
		panic("Division by zero in Pearson calculation.")
	}
	return numerator / denominator
}

func LoadCSV(filePath string) ([]float64, []float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read CSV file: %v", err)
	}

	var test1, test2 []float64
	for _, record := range records[1:] {
		val1, err1 := strconv.ParseFloat(record[0], 64)
		val2, err2 := strconv.ParseFloat(record[1], 64)
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("failed to parse float values: %v, %v", err1, err2)
		}
		test1 = append(test1, val1)
		test2 = append(test2, val2)
	}
	return test1, test2, nil
}

func main() {

	filePath := "final_result.csv"

	test1, test2, err := LoadCSV(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	pearson := PearsonCorrelation(test1, test2)
	fmt.Printf("Pearson Correlation: %.2f\n", pearson)
}
