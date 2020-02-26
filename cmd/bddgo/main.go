package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"plugin"
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

func load(packageName string, functionName string) (
	handler http.Handler,
	err error,
) {

	p, err := plugin.Open(fmt.Sprintf("%s.so", packageName))
	if err != nil {
		return
	}

	sym, err := p.Lookup(functionName)
	if err != nil {
		return
	}

	handler = sym.(func() http.Handler)()
	return
}

func serveHandler(handler http.Handler) (err error) {
	server := http.Server{
		Handler: handler,
	}

	unixListener, err := net.Listen("unix", ".bddgo.s")
	if err != nil {
		return
	}
	server.Serve(unixListener)
	return
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

	handler, err := load(packageName, functionName)
	if err != nil {
		panic(err)
	}

	err = serveHandler(handler)
	if err != nil {
		panic(err)
	}
}
