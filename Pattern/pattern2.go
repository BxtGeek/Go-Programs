package main

import "fmt"

func main() {
	var i, j, row int
	row = 5
	for i = 1; i <= row; i++ {
		for j = 1; j <= 5; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
