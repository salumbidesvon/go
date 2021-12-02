package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	content := "This needs to go in file - github.com/von-salumbides"

	file, err := os.Create("./mygofile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	log.Println("Length is:", length)
	defer file.Close()
	readFile("mygofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename) // data that is being read is in byte format
	checkNilErr(err)
	fmt.Println("Text data in file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
