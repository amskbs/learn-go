package main

import (
	"fmt"
	tp "learn-go/toppackage"
	"learn-go/toppackage/middlepackage/bottompackage"
	"time"
)

type Age int

func (age Age) prt() {
	fmt.Println("I am", age)
}

type Person struct {
	Name        string
	Email       string
	dateOfBirth time.Time
}

func main() {
	var age = Age(15)
	age.prt()
	tp.Test()
	fmt.Println(bottompackage.AddInts(2, 4))
	fmt.Println("Well done!")

	defer metricTime(time.Now())
	a := "some text"
	defer func(s string) {
		fmt.Println(s)
	}(a)
	a = "another text"
	fmt.Println(unintuitive())
	fmt.Println(intuitive())
	man := Person{
		Name:  "Alex",
		Email: "alex@yandex.ru",
	}
	fmt.Println(man)
	m := map[int]string{1: "first"}
	rt := map[int]int{1: 56, 4: 34, 6: 1231}
	fmt.Println("rt = ", rt)
	v, ok := m[1]
	fmt.Println(v, ok)
	delete(m, 2) // ошибки не будет
	delete(m, 1)
	v, ok = m[1]
	fmt.Println(v, ok)
	main2()
	var dest []int
	dest2, dest3 := make([]int, 3), make([]int, 5)
	src := []int{1, 2, 3, 4}
	copy(dest, src)
	copy(dest2, src)
	copy(dest3, src)
	fmt.Println(dest, dest2, dest3, src) // [] [1 2 3] [1 2 3 4 0] [1 2 3 4]

	s := []int{1, 2, 3, 4, 5}
	i := 2

	if len(s) != 0 && i < len(s) { // защищаемся от паники
		s = append(s[:i], s[i+1:]...)
	}
	fmt.Println(s) // [1 2 4 5]
}

func metricTime(start time.Time) {
	// функция Now() возвращает текущее время, а функция Sub возвращает разницу между двумя временными метками
	fmt.Println(time.Since(start))
}

// not clear!!!
func unintuitive() (value string) {
	defer func() { value = "На самом деле" }() // круглые скобки в конце означают, что функция вызывается
	return "Казалось бы"
}

// not clear!!!
func intuitive() string {
	value := "Казалось бы"
	defer func() { value = "На самом деле" }()
	return value
}

func main2() {
	var slice = make([]int, 100)

	for index := range slice {
		slice[index] = index
	}

	fmt.Println(slice)

	slice = slice[:12]

	for i := 0; i < len(slice)/2+1; i++ {
		fmt.Println(i, len(slice)-i-1)
		slice[i], slice[len(slice)-i-1] = slice[len(slice)-i-1], slice[i]
	}

	// slice[0], slice[9] = slice[9], slice[0]

	fmt.Println(slice)
	// a := 1
	// p := &a
	// fmt.Println(p)
	// b := &p
	// fmt.Println(b)

	// *p = 3
	// fmt.Println(a)
	// **b = 5
	// fmt.Println(a)

	// a += 4 + *p + **b
	// fmt.Printf("%d", *p)
}

func test() string {
	return "hi"
}
