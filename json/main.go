package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	First_name string `json:"First_name"`
	Last_name  string `json:"Last_name"`
	Age        int    `json:"Age"`
}

func main() {
	myJson := `
	[
		{
			"First_name": "Von",
			"Last_name": "Salumbides",
			"Age": 25
		},
		{
			"First_name": "Lalisa",
			"Last_name": "Manobal",
			"Age": 24
		}
	]
	`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("unmarshalled: %v", unmarshalled)

	// write json from struct

	var mySlice []Person

	var m1 Person
	m1.First_name = "Wally"
	m1.Last_name = "West"
	m1.Age = 20

	var m2 Person
	m1.First_name = "Diane"
	m1.Last_name = "Croods"
	m1.Age = 21

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("Error marshalling json", err)
	}
	fmt.Print(string(newJson))
}
