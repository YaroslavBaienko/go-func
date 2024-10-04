package main

import (
	"errors"
	"fmt"
	"math"
)

// CalculateVariance concurrently calculates the variance and sends the result to a channel.
func CalculateVariance(numbers []int) (*float64, error) {
	if numbers == nil || len(numbers) == 0 {
		return nil, errors.New("empty number")
	}
	n := float64(len(numbers))
	sum := func(f func(int) float64) float64 {
		var s float64
		results := make(chan float64, len(numbers))
		for _, num := range numbers {
			go func(i int) {
				results <- f(i)
			}(num)
		}
		for i := 0; i < len(numbers); i++ {
			s += <-results
		}
		return s
	}
	mean := sum(func(i int) float64 {
		return float64(i)
	}) / n
	sumRes := sum(func(i int) float64 {
		return math.Pow(float64(i)-mean, 2)
	})
	variance := sumRes / n
	return &variance, nil
}

func main() {
	// Test input
	numbers := []int{5, 455, 34, 36, 94, 4, 49}

	// Calculate variance
	result, err := CalculateVariance(numbers)

	// Check for errors and print the result
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Variance: %.6f\n", *result)
	}

}
