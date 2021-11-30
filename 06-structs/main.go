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

	fmt.Printf("Lisa's details are %+v \n", lisa)
	fmt.Printf("First name is %v, and contact info is %v\n", lisa.firstname, lisa.contactInfo)

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
	fmt.Printf("%+v\n", p)
}
