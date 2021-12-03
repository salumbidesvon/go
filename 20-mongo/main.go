package main

import (
	"fmt"
	"net/http"

	"github.com/von-salumbides/go/20-mongo/router"
)

func main() {
	fmt.Println("Mongo API")
	r := router.Router()

	fmt.Println("Server is getting started...")
	http.ListenAndServe(":4000", r)
	fmt.Println("Listening at port 4000...")
}
