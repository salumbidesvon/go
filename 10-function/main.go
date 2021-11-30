package main

import "fmt"

func main() {
	fmt.Println(adder(1, 2))

	proRes, message := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Result is", proRes)
	fmt.Println("Pro message", proRes, message)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Result function"
}
