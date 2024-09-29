package main

import (
	"errors"
	"fmt"
)

// Функция для вычисления квадрата числа
func Sq(x float64) float64 {
	return x * x
}

// Функция для вычисления суммы элементов с использованием переданной функции
func Sum(xs []float64, f func(float64) float64) float64 {
	sum := 0.0
	for _, x := range xs {
		sum += f(x)
	}
	return sum
}

// Функция для вычисления среднего значения
func Mean(xs []float64) float64 {
	n := float64(len(xs))
	return Sum(xs, func(x float64) float64 { return x }) / n
}

// Функция для вычисления суммы квадратов отклонений от среднего
func SumRes(xs []float64, mean float64) float64 {
	return Sum(xs, func(x float64) float64 { return Sq(x - mean) })
}

// Функция для вычисления дисперсии
func Variance(xs []float64) (float64, error) {
	n := len(xs)
	if n == 0 {
		return 0, errors.New("slice is empty")
	}
	mean := Mean(xs)
	sumRes := SumRes(xs, mean)
	return sumRes / float64(n), nil
}

func main() {
	data := []float64{4, 6, 9, 45, 8, 17, 3}
	variance, err := Variance(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Variance: %.6f\n", variance)
}
