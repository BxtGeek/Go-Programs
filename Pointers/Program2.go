package main

import "fmt"

func main() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Println("Address of Variable A :", &a)
	fmt.Println("Address stored in IP Variable:", ip)
	fmt.Println("Value of *ip variable:", *ip)
}
