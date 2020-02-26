package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Command func([]string)

func parseArguments() (command Command, err error) {
	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage: %s {init|serve|test}\n",
			os.Args[0],
		)
		flag.PrintDefaults()
	}

	if len(os.Args) <= 1 {
		err = errors.New("Insufficient arguments")
		return
	}

	switch os.Args[1] {
	case "init":
		command = InitializeCommand

	case "serve":
		command = ServeCommand

	default:
		err = errors.New("Invalid subcommand")
		return
	}

	return
}

func main() {
	command, err := parseArguments()
	if err != nil {
		flag.Usage()
		os.Exit(1)
	}

	command(os.Args[2:])
}
