package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Learning time in go")

	presentTime := time.Now()
	fmt.Println("Time is", presentTime)

	fmt.Println("Time is", presentTime.Format("01-02-2006 15:02 Monday"))
	log.Println("Time")
}
