package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Multiple types")
	numbers := []uint64{5, 45, 45, 23, 944}
	result, err := CalculateMean(numbers)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Mean: %.7f\n", *result)
	}
}

func CalculateMean[T any](numbers []T) (*float64, error) {
	if len(numbers) == 0 {
		return nil, errors.New("slice is empty")
	}

	var sum uint64
	for _, number := range numbers {
		// Приводим значение к uint64, так как у нас тип uint64
		switch v := any(number).(type) {
		case int:
			sum += uint64(v)
		case uint64:
			sum += v
		default:
			return nil, errors.New("unsupported type")
		}
	}

	// Приводим сумму к float64 перед делением
	n := float64(len(numbers))
	mean := float64(sum) / n
	return &mean, nil
}
