package main

import "fmt"

func main() {
	var a int = 10

LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Println("Value of A is :", a)
		a++
	}
}
