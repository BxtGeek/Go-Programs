// Array of pointers
package main

import "fmt"

func main() {
	a := []int{10, 100, 200}
	var i int

	for i = 0; i < 3; i++ {
		fmt.Printf("Value of a[%d]=%d\n", i, a[i])
	}
}
