package main

//Call by value
//It shows that there is no change in the values though they had been changed inside the function.
import "fmt"

func main() {
	var a int = 10
	var b int = 20
	fmt.Println("Before SWAP value of A:", a)
	fmt.Println("Before SWAP value of B:", b)
	swap(a, b)
	fmt.Println("After SWAP value of A:", a)
	fmt.Println("After SWAP value of B:", b)
}
func swap(x, y int) int {
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}
