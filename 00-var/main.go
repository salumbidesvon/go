package main

import (
	"fmt"
)

// global := "This var is not allowed"
const constant string = "This value is constant"

var GlobalVar string = "This var is global"

func main() {
	var username string = "von"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type: %T \n", smallInt)

	var smallFloat float32 = 255.12345678912345
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var largeFloat float64 = 255.12345678912345
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n", largeFloat)

	// default value and some aliases
	var name string
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)

	anotherVar := "New Var"
	fmt.Println(anotherVar)
	fmt.Println(constant)
	fmt.Println(GlobalVar)

}
