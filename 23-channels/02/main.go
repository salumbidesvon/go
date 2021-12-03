package main

import (
	"log"
	"math/rand"
	"time"
)

const numPool = 10

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}

func CalculateValue(intChan chan int) {
	randmonNumber := RandomNumber(numPool)
	intChan <- randmonNumber
}

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}
