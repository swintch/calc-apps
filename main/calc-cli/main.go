package main

import (
	"flag"
	"log"
	"os"

	"github.com/swintch/calc-apps/handlers"
)

func main() {
	var operator string
	flag.StringVar(&operator, "op", "+", "the operator to use")
	flag.Parse()
	handle := handlers.NewCLIHandler(operator, os.Stdout)
	err := handle.Handler(flag.Args())
	if err != nil {
		log.Fatalln(err)
	}
}
