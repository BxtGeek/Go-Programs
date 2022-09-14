package main

//Call by reference
import "fmt"

func main() {
	var a int = 10
	var b int = 20
	fmt.Println("Before SWAP value of A:", a)
	fmt.Println("Before SWAP value of B", b)
	swap(&a, &b)
	fmt.Println("After SWAP value of A:", a)
	fmt.Println("After SWAP value of B:", b)
}
func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
