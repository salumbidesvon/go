package main

import (
	"fmt"
	"log"
)

func main() {
	for i := 0; i <= 10; i++ {
		log.Println("i is set to", i)

		if i == 3 {
			goto lco
		}
	}
lco:
	fmt.Println("Print me")

	animals := []string{"dog", "cat", "fish", "bird"}

	for _, animal := range animals {
		log.Println("animal is", animal)
	}

	cars := map[string]string{
		"Ford":   "Raptor",
		"Toyota": "Hilux",
		"Nissan": "Navarra",
	}

	for brand, car := range cars {
		log.Println(brand, car)
	}

	type User struct {
		FirstName string
		Lastname  string
		Age       int
	}

	var users []User
	users = append(users, User{"Lalisa", "Manobal", 24})
	users = append(users, User{"Von", "Jethro", 25})

	for _, l := range users {
		log.Println(l, l.FirstName, l.Lastname, l.Age)
	}
}
