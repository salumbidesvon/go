package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One routing")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done() //pass when it's done
	}(wg, mut)
	//wg.Add(1)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two routing")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three routing")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
