package main

import "fmt"

func main() {
	var a int = 10
	for a < 20 {
		if a == 19 {
			fmt.Println("Value of A is :", a)
			a = a + 1
			fmt.Println("Value of A is :", a)
			continue
		}
		fmt.Println("Value of A is :", a)
		a++
	}
}
