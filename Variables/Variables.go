package main

import "fmt"

func main() {
	/*
		// declaring variable in one way
		var a int = 6
		var b int = 5
		var c int
		c = a + b
		fmt.Println(c)
		fmt.Printf("a is of type %T\n", a)
		fmt.Printf("b is of type %T\n", b)
		fmt.Printf("c is of type %T\n", c)
		fmt.Println("*************************************")
	*/

	/*
		// declaring variable in second way
		var q, w int
		q = 1
		w = 2
		e := q + w
		fmt.Println(e)
		fmt.Println("*************************************")
	*/

	// declaring variable in Third way
	var x, y, z = 1, 2, "manoj"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("x is a type of %T\n", x)
	fmt.Printf("y is a type of %T\n", y)
	fmt.Printf("z is a type of %T\n", z)
	fmt.Println("*************************************")

}
