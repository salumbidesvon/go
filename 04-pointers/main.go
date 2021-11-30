package main

import "fmt"

func main() {
	var ptr *int

	fmt.Println("Value of ptr is", ptr)
	fmt.Printf("This is of type %T \n", ptr)

	myNum := 23

	numPtr := &myNum
	fmt.Println("Value of actual pointer is", numPtr)
	fmt.Println("Value of actual pointer is", *numPtr)
}

// package main

// import "log"

// func main() {
// 	myString := "Green"
// 	log.Println("String set to", myString)

// 	changeUsingPointer(&myString)
// 	log.Println("String is now set to", myString)
// }

// func changeUsingPointer(s *string) {
// 	log.Println("s is set to", s)
// 	newValue := "Red"
// 	*s = newValue
// }
