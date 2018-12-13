package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/news/src/router"
)

func main() {
	r := router.NewRouter()
	fmt.Println("Weather Widget")
	log.Fatal(http.ListenAndServe(":8300", r))
}
