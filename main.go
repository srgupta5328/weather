package main

import (
	"log"
	"net/http"

	"github.com/srgupta5328/weather/router"
)

func main() {
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8300", r))
}
