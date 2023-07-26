package main

import (
	"flag"
	"fmt"
	"log-parser/arguments"
	"log-parser/internal"
	"os"
)

func main() {
	args := arguments.New()
	if err := args.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		os.Exit(2)
	}
	internal.Start(args)
}
