package main

import "fmt"

var grade string = "B"
var mark int = 80

func main() {

	switch mark {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("Excellent")
	case grade == "B":
		fmt.Println("Well Done")
	case grade == "c":
		fmt.Println("You Passed")
	default:
		fmt.Println("Failed")
	}
	fmt.Println("your grade is", grade)
}
