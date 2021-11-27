package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "ff000",
		"green": "7455",
		"white": "fffff",
	}
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
