package main

import (
	"fmt"
)

// Person — структура, описывающая человека.
type Person struct {
	Name string
	Year int
}

// NewPerson возвращает новую структуру Person.
func NewPerson(name string, year int) Person {
	return Person{
		Name: name,
		Year: year,
	}
}

// String возвращает информацию о человеке.
func (p Person) String() string {
	return fmt.Sprintf("Имя: %s, Год рождения: %d", p.Name, p.Year)
}

// Print выводит информацию о человеке.
func (p Person) Print() {
	// вызовется метод String() для Person
	fmt.Println(p)
}

// Person — структура, описывающая человека.
// type Person2 struct {
// 	Name string
// 	Year int
// }

// // NewPerson возвращает новую структуру Person.
// func NewPerson2(name string, year int) Person2 {
// 	return Person2{
// 		Name: name,
// 		Year: year,
// 	}
// }

// // String возвращает информацию о человеке.
// func (p Person2) String() string {
// 	return fmt.Sprintf("Имя: %s, Год рождения: %d", p.Name, p.Year)
// }

// // Print выводит информацию о человеке.
// func (p Person2) Print() {
// 	// вызовется метод String() для Person
// 	fmt.Println(p)
// }

// Student описывает студента с использованием вложенной структуры Person. То есть структура Student описывает.
type Student struct {
	Person  // вложенный объект Person
	Person2 // вложенный объект Person2
	Group   string
}

func NewStudent(name string, year int, group string) Student {
	return Student{
		Person: NewPerson(name, year), // Явно создаём структуру Person
		Group:  group,
	}
}

// String возвращает информацию о студенте.
func (s Student) String() string {
	return fmt.Sprintf("%s, Группа: %s", s.Person, s.Group)
}

func main() {
	s := NewStudent("John Doe", 1980, "701")
	s.Print()
	// вызовется метод String() для Student
	fmt.Println(s)
	fmt.Println(s.Person.Name, s.Person.Year, s.Group)
}
