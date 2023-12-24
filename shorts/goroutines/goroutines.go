package main

import (
	"fmt"
	"runtime"
	"strconv"
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

func Printf(v interface{}) {
	switch v2 := v.(type) {
	case int:
		fmt.Print("Это число " + strconv.FormatInt(1, 10))
	case string:
		fmt.Print("Это строка " + v2)
	default:
		fmt.Print("Неизвестный тип")
	}
}
