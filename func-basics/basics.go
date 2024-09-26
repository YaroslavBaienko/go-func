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
		fmt.Printf("ERROR: %v", learnerError)
	} else {
		fmt.Println(learnerResult)
	}
}

// CalculateMean returns mean of the numbers slice, or an error.
func CalculateMean(numbers []int) (*float64, error) {
	// Your code goes here.
	return nil, errors.New("not implemented")
}
