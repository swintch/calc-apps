package main

import (
	"log"
	"os"

	"github.com/swintch/calc-apps/handlers"
	"github.com/swintch/mdw-smarty-calc-lib2/calc"
)

func main() {

	handler := handlers.NewCSVHandler(os.Stdin, os.Stdout, os.Stderr, calculators)
	err := handler.Handle()
	if err != nil {
		log.Fatal(err)
	}
}

var calculators = map[string]calc.Calculator{
	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
}
