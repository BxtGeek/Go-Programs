// Declare an Array Variable with a Specific Size
package main

var a = [5]int{5, 6, 7, 8, 9}
var i, sum int

func main() {
	for i = 0; i < 5; i++ {
		println("Value of i = ", a[i])
	}
}
