package main

import "fmt"

// Create new type of 'Deck'
// which is a slice of string
type deck []string

func (d deck) print() {
	for i, card := range d { // To iterate
		fmt.Println(i, card)
	}
}
