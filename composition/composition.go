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

// Функция для вычисления дисперсии с преобразованием []int в []float64
func CalculateVariance(xs []int) (*float64, error) {
	n := len(xs)
	if n == 0 {
		return nil, errors.New("slice is empty")
	}

	// Преобразование []int в []float64
	floats := make([]float64, n)
	for i, v := range xs {
		floats[i] = float64(v)
	}

	// Вычисление среднего
	mean := Mean(floats)

	// Вычисление суммы квадратов отклонений
	sumRes := SumRes(floats, mean)

	// Вычисляем дисперсию
	variance := sumRes / float64(n)

	// Возвращаем указатель на результат
	return &variance, nil
}

// Основная функция для тестирования
func main() {
	// Пример входных данных
	numbers := []int{4, 6, 9, 45, 8, 17, 3}

	// Вычисляем дисперсию
	learnerResult, learnerError := CalculateVariance(numbers)
	if learnerError != nil {
		fmt.Println("Error:", learnerError)
		return
	}

	// Печатаем результат
	fmt.Printf("Variance: %.6f\n", *learnerResult)
}
