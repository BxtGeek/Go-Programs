package main

import "fmt"

var n [10]int
var i int

func main() {
	for i = 0; i < 5; i++ {
		n[i] = i + 100
		fmt.Println("Current value is:", n[i])
	}
}
