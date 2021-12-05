package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in go")

	myChan := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myChan <- 5
	// fmt.Println(<-myChan)
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChOpen := <-myChan

		fmt.Println(isChOpen)
		fmt.Println(val)
		fmt.Println(<-myChan)

		wg.Done()
	}(myChan, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChan <- 5
		myChan <- 6
		close(myChan)
		wg.Done()
	}(myChan, wg)

	wg.Wait()
}
