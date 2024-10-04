package main

import (
	"errors"
	"slices"
)

func Median(numbers []int) (*float64, error) {
	if len(numbers) == 0 {
		return nil, errors.New("slice is empty")
	}

	slices.Sort(numbers)

	var median float64

	var isOdd bool

	if len(numbers)%2 == 0 {
		isOdd = true
	} else {
		isOdd = false
	}

	if isOdd {
		middle1 := numbers[len(numbers)/2-1]
		middle2 := numbers[len(numbers)/2]
		median = (float64(middle1) + float64(middle2)) / 2.0
	} else {
		median = float64(numbers[len(numbers)/2])
	}
	return &median, nil

}
