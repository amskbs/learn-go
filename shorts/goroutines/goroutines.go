package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("yeah baby 1!")
	runtime.GOMAXPROCS(1)

	start := time.Now()
	go func() {
		fmt.Println("yeah baby 2!", time.Since(start))
	}()

	for {
	}
}
