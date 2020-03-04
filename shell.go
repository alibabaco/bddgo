package bddgo

import (
	"fmt"
	"os/exec"
	"strings"
)

func shell(format string, a ...interface{}) error {
	commandString := fmt.Sprintf(format, a...)
	parts := strings.Split(commandString, " ")
	//var arguments [len(a) - 1]string
	//copy(arguments[:], parts[1:])
	cmd := exec.Command(parts[0], parts[1:]...)
	fmt.Println(cmd)
	return nil
}
