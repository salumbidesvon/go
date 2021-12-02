package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // to change the vale from Name -> coursename
	Price    int
	Platform string
	Password string   `json:"-"`              // hide the password
	Tags     []string `json:"tags,omitempty"` // removed the nil value
}

func main() {
	fmt.Println("Learning Json...")
	// encodeJson()
	decodeJson()
}

func encodeJson() {
	courses := []course{
		{"Go", 255, "MyGoLang", "abc1", []string{"back-dev", "go"}},
		{"Node", 254, "MyNode", "abc2", []string{"front-dev", "node"}},
		{"Python", 256, "MyPy", "abc3", nil},
	}

	// package data as JSON data

	// finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonData := []byte(`
	{
			"coursename": "Go",
			"Price": 255,
			"Platform": "MyGoLang",
			"tags": ["back-dev","go"]
	}
	`)

	var mycourse course

	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &mycourse)
		fmt.Printf("%#v\n", mycourse)
	} else {
		fmt.Println("The JSON was not valid")
	}

	// some cases you want to add data
	var myData map[string]interface{} // Use interface if you don't know the value of data
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key is %v and value is %v and type is %T \n", k, v, v)
	}
}
