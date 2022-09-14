package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string = "Manoj"
	var b string = "Manoj"
	var result int
	result = comp(a, b)
	if result == 0 {
		fmt.Println("Both String are same")
	} else {
		fmt.Println("Both String are not same")
	}
}
func comp(x, y string) int {
	var result int
	result = (strings.Compare(x, y))
	return result
}
