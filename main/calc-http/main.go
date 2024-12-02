package main

import (
	"log"
	"net/http"

	"github.com/swintch/calc-apps/handlers"
)

func main() {
	err := http.ListenAndServe("localhost:8080", handlers.NewHTTPRouter())
	if err != nil {
		log.Fatal(err)
	}
}
