package main

import "fmt"

func main() {
	var i, j, row int
	row = 5
	for i = 1; i <= row; i++ {
		for j = 1; j <= row-i+1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}
