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
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [packages]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(functionName, "function", "GetMainHandler",
		"Function name to load and call to get the main handler.")
	flag.Parse()

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
