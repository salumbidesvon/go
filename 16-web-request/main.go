package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web request")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	performPostFormRequest()
}

func performPostFormRequest() {
	const myurl = "https://webhook.site/9f9b27bf-91ae-4f8a-b260-3fee4e7d18b8"

	// form data
	data := url.Values{}
	data.Add("Firstname", "von")
	data.Add("lastname", "vsalumbides")
	data.Add("language", "go")

	response, err := http.PostForm(myurl, data)
	checkNilErr(err)
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println(string(content))
}

func PerformGetRequest() {
	const myrul = "https://webhook.site/9f9b27bf-91ae-4f8a-b260-3fee4e7d18b8"

	response, err := http.Get(myrul)
	checkNilErr(err)
	defer response.Body.Close()

	fmt.Println("Status code:", response.Status)
	fmt.Println("Content Length:", response.ContentLength)

	// fmt.Println("Content is:", string(content))

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Bytecount is:", byteCount)
	fmt.Println(responseString.String())

}

func PerformPostJsonRequest() {
	const myurl = "https://webhook.site/9f9b27bf-91ae-4f8a-b260-3fee4e7d18b8"

	// Fake json payload

	requestBody := strings.NewReader(`
		{
			"firstname" : "von",
			"lastname" : "salumbides",
			"language" : "go"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	checkNilErr(err)
	defer response.Body.Close()
	// log.Println("Response:", response)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	// var responseString strings.Builder
	// byteCount, _ := responseString.Write(content)
	// fmt.Println("Bytecount is:", byteCount)
	// fmt.Println(responseString.String())
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
