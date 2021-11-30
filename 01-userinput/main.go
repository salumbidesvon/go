package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your favorite word")

	// comma ok // err ok
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("That's awesome ", input)
	fmt.Printf("Type of input is %T \n", input)

}
