package main

import (
	"errors"
	"fmt"
	"math"
)

// CalculateVariance returns variance of the numbers slice, or an error.
func CalculateVariance(numbers []int) (*float64, error) {
	// Проверяем, что слайс не пустой. Если пустой, возвращаем ошибку.
	if numbers == nil || len(numbers) == 0 {
		return nil, errors.New("empty numbers")
	}

	// Преобразуем длину массива в тип float64, так как нам нужно работать с плавающей точкой.
	n := float64(len(numbers))

	// Функция sum принимает функцию f как параметр и возвращает сумму значений слайса numbers.
	sum := func(f func(int) float64) float64 {
		var s float64
		for _, num := range numbers {
			s += f(num)
		}
		return s
	}

	// Вычисляем среднее значение. Для этого суммируем все значения numbers и делим на n.
	mean := sum(func(i int) float64 {
		return float64(i)
	}) / n

	// Вычисляем сумму квадратов отклонений от среднего значения.
	sumRes := sum(func(i int) float64 {
		return math.Pow(float64(i)-mean, 2)
	})

	// Вычисляем дисперсию: сумма отклонений / количество элементов.
	variance := sumRes / n

	// Возвращаем указатель на дисперсию и nil (без ошибки).
	return &variance, nil
}

// Основная функция для тестирования
func main() {
	// Пример входных данных
	numbers := []int{5, 455, 34, 36, 94, 4}

	// Вычисляем дисперсию
	learnerResult, learnerError := CalculateVariance(numbers)
	if learnerError != nil {
		fmt.Println("Error:", learnerError)
		return
	}

	// Печатаем результат
	fmt.Printf("Variance: %.6f\n", *learnerResult)
}
