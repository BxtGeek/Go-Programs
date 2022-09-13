package main

import "fmt"

func main() {
	var a int = 10
	for a < 20 {
		fmt.Println("Value of A is:", a)
		a++
		if a > 15 {
			break
		}
	}
}
