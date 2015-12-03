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
	case "server":
		fmt.Println("Running server...")
		blockRun("rfk-server", args[1:])
	case "skip", "reward", "play", "pause":
		blockRun("rfk-cli", args[0:0])
	default:
		fmt.Printf("Unknown command %q\n", args)
	}
}

func blockRun(cmd string, args []string) {
	foundCmd, err := exec.LookPath("rfk-server")
	panicErr(err)
	fmt.Printf("running: %q\n", foundCmd)

	//// not a fork!
	//command := exec.Command(foundCmd, args...)
	//err = command.Run()
	//panicErr(err)
	//command.Wait()

	_, err = syscall.ForkExec(foundCmd, args, nil)
	panicErr(err)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
