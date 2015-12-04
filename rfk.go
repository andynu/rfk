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
		execRun("rfk-server", args[1:])
	case "skip", "reward", "play", "pause":
		execRun("rfk-cli", args[0:0])
	default:
		fmt.Printf("Unknown command %q\n", args)
	}
}

func execRun(cmd string, args []string) {
	binary, err := exec.LookPath("rfk-server")
	panicErr(err)
	fmt.Printf("running: %q\n", binary)

  var argsWithBinary []string
  argsWithBinary = append(argsWithBinary, binary)
  for _, arg := range args {
    argsWithBinary = append(argsWithBinary, arg)
  }

  env := os.Environ()
	err = syscall.Exec(binary, argsWithBinary, env)
	panicErr(err)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
