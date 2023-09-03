package main

import (
	"fmt"
	"strings"
)

func main() {
	tower := buildTower(10)
	for _, row := range tower {
		fmt.Println(row)
	}
}

func buildTower(nFloors int) []string {
	tower := make([]string, nFloors)
	var count = 1
	for i := nFloors - 1; i >= 0; i-- {
		tower[nFloors-1-i] = strings.Repeat(" ", i) + strings.Repeat("*", count) + strings.Repeat(" ", i)
		count += 2
	}
	return tower
}
