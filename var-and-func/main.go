package main

import (
	"log"
)

var s = "seven"

func main() {
	// var firstname string
	s2 := "six"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("five")

	// firstname = "Von"
	// fmt.Println(firstname)

	// whatWasSaid, theOtherWhatIsSaid := saySomething()
	// fmt.Println("The function said", whatWasSaid, theOtherWhatIsSaid)
}

func saySomething(s string) (string, string) {
	log.Println("Log from s is", s)
	return s, "there"
}
