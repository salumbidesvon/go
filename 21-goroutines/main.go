package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

// Use case : for databases with replica
func main() {
	// go greeter("Hello")
	// greeter("World")

	websites := []string{
		"http://google.com",
		"http://fb.com",
		"http://github.com",
		"http://amazon.com",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("Signals", signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status for %s\n", res.StatusCode, endpoint)
	}
}
