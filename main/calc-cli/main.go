package main

import (
	"log"
	"os"

	"github.com/swintch/calc-apps/handlers"
	"github.com/swintch/mdw-smarty-calc-lib2/calc"
)

func main() {
	calculator := &calc.Addition{}
	handlerObj := handlers.NewCLIHandler(calculator, os.Stdout)
	err := handlerObj.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

}
