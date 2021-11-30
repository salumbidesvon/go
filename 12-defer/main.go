package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")

	fmt.Println("Hello")
	myDefer() // this will call first because the defer was called inside the function
}

func myDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("I is: ", i)
	}
}
