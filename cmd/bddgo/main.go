package main

import (
	"flag"
	"fmt"
	"github.com/alibabaco/bddgo"
	"os"
	"path/filepath"
)

func parseArguments(packageName *string, functionName *string) {
	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage: %s {init|test|load}\n",
			os.Args[0],
		)
		flag.PrintDefaults()
	}

	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	loadCommand := flag.NewFlagSet("load", flag.ExitOnError)
	loadCommand.StringVar(functionName, "function", "GetMainHandler",
		"Function name to load and call to get the main handler.")
	loadCommand.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage: %s [packages]\n",
			os.Args[0],
		)
		flag.PrintDefaults()
	}

	if len(os.Args) <= 1 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		initCommand.Parse(os.Args[2:])

	case "load":
		loadCommand.Parse(os.Args[2:])

	default:
		flag.Usage()
		os.Exit(1)
	}

	if flag.NArg() > 0 {
		*packageName = flag.Args()[0]
	}
}

func main() {
	packageName := "."
	var functionName string
	parseArguments(&packageName, &functionName)

	packageName, err := filepath.Abs(packageName)
	if err != nil {
		panic(err)
	}
	packageName = filepath.Base(packageName)

	handler, err := bddgo.LoadPackageBinary(packageName, functionName)
	if err != nil {
		panic(err)
	}

	err = bddgo.ServeHandler(handler, fmt.Sprintf(".%s.s", packageName))
	if err != nil {
		panic(err)
	}
}
