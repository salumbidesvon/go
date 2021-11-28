package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type tagalogBot struct{}

func main() {
	eb := englishBot{}

	tb := tagalogBot{}

	printGreeting(eb)
	printGreeting(tb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (tagalogBot) getGreeting() string {
	return "Kamusta!"
}
