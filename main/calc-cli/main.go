package main

import (
	"log"
	"os"

	"github.com/calc"
	"github.com/calc-apps/handlers"
)

func main() {
	handle := handlers.NewCLIHandler(&calc.Addition{}, os.Stdout)
	err := handle.Handler(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
}
