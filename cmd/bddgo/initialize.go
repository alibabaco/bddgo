package main

import (
	"flag"
)

func InitializeCommand(arguments []string) {
	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	initCommand.Parse(arguments)
}
