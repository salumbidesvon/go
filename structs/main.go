package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	lisa := person{
		firstname: "La",
		lastname:  "Lisa",
		contactInfo: contactInfo{
			email:   "lalisa@gmail.com",
			zipCode: 143,
		},
	}

	lisa.updateName("Lala")
	lisa.print()

	lisaPointer := &lisa
	lisaPointer.updateName("Lala")
	lisa.print()

	lisa.updateName("Lala")
	lisa.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
