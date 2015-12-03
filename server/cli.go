package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/andynu/rfk/server/karma"
	"github.com/andynu/rfk/server/player"
)

func consoleInputListener() {
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, _ := reader.ReadString('\n')
			text = strings.TrimSuffix(text, "\n")
			switch text {
			case "n":
				log.Printf("cmd: n - Stop")
				player.Stop()
			case "s":
				log.Printf("cmd: s - Skip")
				player.Skip()
			case "r":
				log.Printf("cmd: r - Reward")
				karma.Log(player.CurrentSong, 1)
			default:
				log.Printf("cmd: %q - unknown command", text)
			}
		}
	}()
}
