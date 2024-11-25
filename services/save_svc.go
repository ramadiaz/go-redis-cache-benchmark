package services

import (
	"encoding/csv"
	"fmt"
	"os"
)

func SaveResultsToCSV(filename string, keys []string, times []float64) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	
	err = writer.Write([]string{"Key", "Response Time (ms)"})
	if err != nil {
		return err
	}

	
	for i := range keys {
		err = writer.Write([]string{keys[i], fmt.Sprintf("%.2f", times[i])})
		if err != nil {
			return err
		}
	}

	return nil
}

func SaveFinalResultToCSV(filename string, test1, test2 []float64) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	
	err = writer.Write([]string{"test1", "test2"})
	if err != nil {
		return fmt.Errorf("failed to write header: %v", err)
	}

	
	for i := 0; i < len(test1); i++ {
		row := []string{
			fmt.Sprintf("%f", test1[i]),
			fmt.Sprintf("%f", test2[i]),
		}
		err := writer.Write(row)
		if err != nil {
			return fmt.Errorf("failed to write row: %v", err)
		}
	}

	return nil
}