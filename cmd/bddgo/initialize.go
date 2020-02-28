package main

import (
	"flag"
	"fmt"
	"github.com/alibabaco/bddgo"
	"os"
	"os/exec"
)

func InitializeCommand(arguments []string) {
	defaultPython, err := exec.LookPath("python3")
	if err != nil {
		panic("Cannot retrieve default python3 interpreter")
	}
	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	chdir := initCommand.String(
		"C",
		".",
		"Change directory before initialize.",
	)
	python := initCommand.String(
		"p",
		defaultPython,
		"Python interpreter to use",
	)
	initCommand.Parse(arguments)
	initCommand.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage: %s %s\n",
			os.Args[0],
			os.Args[1],
		)
		flag.PrintDefaults()
	}

	//fmt.Printf("%s, %s", *chdir, *python)
	venv := pyvenv.New(chdir, python)
	if venv.Exists() {
		fmt.Printf("Python virtualenv is already exists: %s\n", venv.path)
	}

}
