package main

import "log"

func main() {
	myString := "Green"
	log.Println("String set to", myString)

	changeUsingPointer(&myString)
	log.Println("String is now set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
