package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Commands maps commands to handler functions
var Commands = map[string]func() error{}

func init() {
	flag.Usage = usage
}

func usage() {
	keys := make([]string, 0, len(Commands))
	for k := range Commands {
		keys = append(keys, k)
	}
	fmt.Fprintf(os.Stderr, "Usage: %s [%s] <flags>\n\n", os.Args[0], strings.Join(keys, "|"))
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}
