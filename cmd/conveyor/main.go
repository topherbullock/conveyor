package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

type ConveyorCommand struct {
}

func main() {
	var cmd ConveyorCommand

	parser := flags.NewParser(&cmd, flags.HelpFlag|flags.PassDoubleDash)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
