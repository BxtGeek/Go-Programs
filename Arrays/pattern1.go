package main

import "fmt"

func main() {
	//var a [5][5]int
	var i, j int
	for i = 1; i <= 5; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
}
