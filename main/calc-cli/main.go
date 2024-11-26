package main

import (
	"flag"
	"log"
	"os"

	"github.com/swintch/calc-apps/handlers"
	"github.com/swintch/mdw-smarty-calc-lib2/calc"
)

func main() {
	var operation string
	flag.StringVar(&operation, "op", "+", "Operation to use.")
	flag.Parse()
	calculator, ok := calculators[operation]
	if !ok {
		log.Fatalf("Unknown operation %s.", operation)
	}
	handlerObj := handlers.NewCLIHandler(calculator, os.Stdout)
	err := handlerObj.Handle(flag.Args())
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
