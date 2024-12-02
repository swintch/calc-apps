package main

import (
	"log"
	"os"

	"github.com/swintch/calc-apps/handlers"
)

func main() {
	handle := handlers.NewCSVHandler(os.Stdin, os.Stdout, os.Stdout)
	err := handle.Handle()
	if err != nil {
		log.Fatal(err)
	}
}
