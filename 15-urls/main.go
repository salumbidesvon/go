package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://github.com/von-salumbides?tab=repositories"

func main() {
	fmt.Println("Handling URL in Go")

	result, _ := url.Parse(myUrl)
	// fmt.Println("Result:", result.Scheme)
	// fmt.Println("Result:", result.Host)
	// fmt.Println("Result:", result.Path)
	// fmt.Println("Result:", result.RawQuery)
	// fmt.Println("Result:", result.Port())

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T \n", qparams)
	fmt.Println("Values: ", qparams["repositories"])

	for _, val := range qparams {
		fmt.Println("Params is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   "von-salumbides",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
