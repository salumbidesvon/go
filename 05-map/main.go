package main

import "fmt"

func main() {
	// myMap := make(map[string]string) -> Makes a map

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
