package main

import (
	"log"
	"net/http"

	"github.com/news/src/router"
)

func main() {
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8300", r))
}
