package main

import (
	"flag"
	"log"
	"os"

	"github.com/joerx/grpc_routeguide/cmd"
)

var foo string

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	handler := cmd.Commands[flag.Arg(0)]
	if handler == nil {
		flag.Usage()
		os.Exit(1)
	}
	if err := handler(); err != nil {
		log.Fatal(err)
	}
}
