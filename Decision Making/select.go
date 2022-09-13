package main

import (
	"fmt"
	"time"
)

func main() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		one <- "hey"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		two <- "Hello"
	}()

	select {
	case rec1 := <-one:
		fmt.Println("I recevied from one channel!", rec1)

	case rec2 := <-two:
		fmt.Println("I recevied from second channel!", rec2)
	}
}
