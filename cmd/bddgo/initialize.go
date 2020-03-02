package main

import (
	"flag"
	"fmt"
	"github.com/alibabaco/bddgo"
	"os"
	"os/exec"
)

func InitializeCommand(arguments []string) error {
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
	recreate := initCommand.Bool(
		"r",
		false,
		"Recreate the environment if exists",
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
	venv := bddgo.VirtualEnv(*chdir, *python)
	if venv.Exists() {
		if *recreate {
			fmt.Printf("Deleting %s\n", venv.Path)
			venv.Delete()
		} else {
			return fmt.Errorf("Python virtualenv is already exists: %s\n", venv.Path)
		}
	}

	err = venv.Create()
	if err != nil {
		return err
	}

	return nil
}
