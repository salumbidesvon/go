package main

import "log"

func main() {
	var isTrue bool

	myNum := 100
	isTrue = true

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is", isTrue)
	} else if myNum < 100 && isTrue {
		log.Println("1")
	} else {
		log.Println("2")
	}
}
