package main

import "fmt"

type LogLevel int

const (
	LogLevelError LogLevel = 2*iota + 3
	LogLevelWarning
	LogLevelInfo
)

func main() {
	fmt.Println(LogLevelError, LogLevelWarning, LogLevelInfo)
}
