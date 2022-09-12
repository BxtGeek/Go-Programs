package main

import "fmt"

var a int = 30
var b int = 40

func main() {
	if a == 10 {
		if b == 20 {
			fmt.Println("A value is 10 and B value is 20")
		}
	}
	fmt.Println("Value of A is", a)
	fmt.Println("Value of B is", b)
}
