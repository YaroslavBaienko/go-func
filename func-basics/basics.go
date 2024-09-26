package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello Go!")
	numbers := []int{4, 6, 9, 45, 8, 17, 3}
	learnerResult, learnerError := CalculateMean(numbers)
	if learnerError != nil {
		fmt.Printf("Error: %v\n", learnerError)
	} else {
		fmt.Printf("Result: %.6f\n", *learnerResult)
	}
}

// CalculateMean returns the mean of the numbers slice, or an error.
func CalculateMean(numbers []int) (*float64, error) {
	length := len(numbers)
	if length == 0 {
		return nil, errors.New("empty numbers")
	}
	var sum float64 = 0.0
	for _, value := range numbers {
		sum += float64(value)
	}

	mean := sum / float64(length)
	return &mean, nil
}
