package main

import "fmt"

// Person — структура, описывающая человека.
type Person2 struct {
	Name string
	Year int
}

// NewPerson возвращает новую структуру Person.
func NewPerson2(name string, year int) Person2 {
	return Person2{
		Name: name,
		Year: year,
	}
}

// String возвращает информацию о человеке.
func (p Person2) String() string {
	return fmt.Sprintf("Имя: %s, Год рождения: %d", p.Name, p.Year)
}

// Print выводит информацию о человеке.
func (p Person2) Print() {
	// вызовется метод String() для Person
	fmt.Println(p)
}
