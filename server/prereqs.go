package main

import (
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

func checkPrereqs() {
	ensureExists("mpg123")
}

func ensureExists(executable string) {
	cmd := exec.Command(executable)
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus := exitError.Sys().(syscall.WaitStatus)
			exitStatus := waitStatus.ExitStatus()
			switch exitStatus {
			case 1:
				log.Printf("prereq: %q OK\n", executable)
				return
			default:
				panic(fmt.Errorf("prereq: Unknown exit %q => %d - %q", executable, exitStatus, err))
			}
		}
		panic(fmt.Errorf("prereq: Unknown error for %q => %q", executable, err))
	}
}
