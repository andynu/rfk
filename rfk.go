package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	switch args[0] {
	case "skip", "reward", "play", "pause":
		execRun("rfk-cli", args[0:]...)
	default:
		if isDashRun(args[0]) {
			execRun("rfk-"+args[0], args[1:]...)
		}
		fmt.Printf("Unknown command %q\n", args)
	}
}

func isDashRun(cmd string) bool {
	_, err := exec.LookPath("rfk-server")
	if err != nil {
		return false
	}
	return true
}

func execRun(cmd string, args ...string) {
	binary, err := exec.LookPath("rfk-server")
	panicErr(err)
	fmt.Printf("running: %q\n", binary)

	binargs := append([]string{cmd}, args...)

	env := os.Environ()
	err = syscall.Exec(cmd, binargs, env)
	panicErr(err)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
