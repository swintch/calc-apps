package main

import (
	"log"
	"net/http"

	"github.com/swintch/calc"
	"github.com/swintch/calc-apps/app/calculator"
	router "github.com/swintch/calc-apps/http"
)

func main() {
	appHandler := calculator.NewHandler(
		calc.Addition{},
		calc.Subtraction{},
		calc.Multiplication{},
		calc.Division{},
	)
	endpoint := "localhost:8080"
	log.Println("Listening on", endpoint)
	err := http.ListenAndServe(endpoint, router.SmartyShuttleRouter(appHandler))
	if err != nil {
		log.Fatalln(err)
	}
}
