package main

import (
	"fmt"
)

// Настройка для управления отображением результата и подсказок
const showExpectedResult = false
const showHints = false

const (
	SuccessMsg string = "Operation successful!"
	PanicMsg   string = "Operation panicked!"
)

// Person - это структура, представляющая человека с именем и возрастом.
type Person struct {
	Name string
	Age  int
}

// Persons хранит карту (map) объектов Person.
type Persons struct {
	m map[string]Person
}

// NewPersons инициализирует объект Persons с пустой картой.
func NewPersons() *Persons {
	return &Persons{
		m: make(map[string]Person),
	}
}

// PanickingAdd добавляет нового Person в Persons и вызывает панику,
// если возраст не соответствует допустимым значениям.
func (p *Persons) PanickingAdd(np Person) {
	if np.Age < 0 || np.Age > 100 {
		panic(fmt.Sprintf("add:invalid age %d for name %s", np.Age, np.Name))
	}
	p.m[np.Name] = np
}

// PopulateData должна добавлять данные в Persons, избегая паники.
func PopulateData(data []Person) (result string) {
	result = SuccessMsg // Устанавливаем результат успешной операции по умолчанию
	p := NewPersons()   // Создаем объект Persons

	// defer гарантирует, что этот код выполнится в конце функции,
	// даже если произойдет паника
	defer func() {
		if r := recover(); r != nil {
			// Если произошла паника (р = информация о панике), меняем результат
			result = PanicMsg
		}
	}()

	// Проходим по всем элементам массива data
	for _, np := range data {
		// Пытаемся добавить каждого человека через PanickingAdd
		p.PanickingAdd(np)
	}

	// Возвращаем результат
	return
}

func main() {
	// Пример использования
	data := []Person{
		{"Alice", 25},
		{"Bob", 150}, // Этот вызовет панику из-за недопустимого возраста
		{"Charlie", 30},
	}

	result := PopulateData(data)
	fmt.Println(result) // Ожидается "Operation panicked!"
}
