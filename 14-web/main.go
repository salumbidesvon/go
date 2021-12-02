package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Web request")

	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response type is %T \n", response)

	dataBytes, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	content := string(dataBytes)
	fmt.Println("Body: ", (content))
	defer response.Body.Close() // caller's responsibility to close the connection
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
