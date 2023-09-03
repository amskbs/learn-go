package main

import (
	"fmt"
)

func main() {
	a := 0
	defer fmt.Println(a)
	defer func() {
		a++
		fmt.Println(a)
	}()

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
