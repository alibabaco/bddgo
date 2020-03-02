package main

import (
	"flag"
	"fmt"
	"github.com/alibabaco/bddgo"
	"os"
	"path/filepath"
)

func ServeCommand(arguments []string) error {
	packageName := "."
	var functionName string

	serveCommand := flag.NewFlagSet("serve", flag.ExitOnError)
	serveCommand.StringVar(&functionName, "function", "GetMainHandler",
		"Function name to load and call to get the main handler.")
	serveCommand.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage: %s %s [packages]\n",
			os.Args[0],
			os.Args[1],
		)
		flag.PrintDefaults()
	}

	serveCommand.Parse(arguments)

	if flag.NArg() > 0 {
		packageName = flag.Args()[0]
	}

	packageName, err := filepath.Abs(packageName)
	if err != nil {
		return err
	}
	packageName = filepath.Base(packageName)

	handler, err := bddgo.LoadPackageBinary(packageName, functionName)
	if err != nil {
		return err
	}

	err = bddgo.ServeHandler(handler, fmt.Sprintf(".%s.s", packageName))
	if err != nil {
		return err
	}
	return nil
}
