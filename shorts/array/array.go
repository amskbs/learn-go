package main

import (
	"fmt"
)

func main() {
	array := [2]int{2, 3}
	fmt.Println(array)
	array[0] = 5
	array[1] = 6
	fmt.Println(array)
}
