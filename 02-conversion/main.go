package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to my home")
	fmt.Println("Please rate my home")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Thanks for rating ", input)

	// numRating store the value of converted input
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error: ", err)
		// panic(err) // cause the program to exit
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}
