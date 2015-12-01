package main

import (
	"bufio"
	"fmt"
	"github.com/andynu/rfk/server/player"
	"os"
	"strings"
)

func listenForInput() {
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, _ := reader.ReadString('\n')
			text = strings.TrimSuffix(text, "\n")
			fmt.Printf("got %q\n", text)
			switch text {
			case "n":
				player.Stop()
			}
		}
	}()
}
