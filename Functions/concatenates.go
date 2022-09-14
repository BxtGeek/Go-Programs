package main

import "fmt"

func main() {
	var a string = "Manoj"
	var b string = "Bxtgeek"
	var result string
	result = concate(a, b)
	fmt.Println("Concate message is :", result)
}

func concate(x, y string) string {
	var result1 string
	result1 = x + " " + y
	return result1
}
