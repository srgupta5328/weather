package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/srgupta5328/weather/api"
)

func main() {
	fmt.Println("Initializing router...")
	r := api.NewRouter()
	fmt.Println("Listening on http://localhost:8300")
	log.Fatal(http.ListenAndServe(":8300", r))
}
