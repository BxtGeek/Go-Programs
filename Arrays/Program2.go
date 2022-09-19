package main

import "fmt"

var arr = [5]string{"Manoj", "Parul", "Arun", "Utkarsk", "Shantanu"}
var i int

func main() {
	for i = 0; i < 5; i++ {
		fmt.Println("Value of array is:", arr[i])
	}
}
