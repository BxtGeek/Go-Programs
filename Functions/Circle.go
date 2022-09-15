package main

import (
	"fmt"
	"math"
)

func main() {
	var radius int = 5
	var area float32
	area = math.Pi * (float32(radius) * float32(radius))
	fmt.Println("Area of a Circle is:", area)
}
