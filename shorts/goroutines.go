package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		// fmt.Println("shceduled goroutine ", i)
		go fmt.Println(i)
	}
	time.Sleep(100000)
}
