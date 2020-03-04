package bddgo

import (
	"fmt"
	"os/exec"
)

func execute(executable string, arguments ...string) error {
	cmd := exec.Command(executable, arguments...)
	err := cmd.Run()
	if err != nil {
		return err
	}

	exitStatus := cmd.ProcessState.ExitCode()
	if exitStatus != 0 {
		return fmt.Errorf(
			"Command %s is failed with exit code: %s\n",
			executable,
			exitStatus,
		)
	}
	return nil
}
