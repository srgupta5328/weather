package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/srgupta5328/weather/router"
)

func main() {
	fmt.Println("Initializing router...")
	r := router.NewRouter()
	fmt.Println("Listening on http://localhost:8300")
	log.Fatal(http.ListenAndServe(":8300", r))
}
