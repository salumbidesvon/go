package main

import "log"

type myStruct struct {
	Firstname string
}

func (m *myStruct) printFirstName() string {
	return m.Firstname
}

func main() {
	var myVar myStruct
	myVar.Firstname = "Lalisa"

	myVar2 := myStruct{
		Firstname: "Jisoo",
	}

	log.Println("My var is set to", myVar.printFirstName())
	log.Println("My var2 is set to", myVar2.printFirstName())
}
