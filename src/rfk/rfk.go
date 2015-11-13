// RFK Entry point, and main loop
package main

import (
	"flag"
	"log"
	"rfk/dj"
	"rfk/player"
)

func main() {
	command := flag.String("c", "", "command")
	flag.Parse()

	switch *command {
	default:
		log.Println("RFK v3 startup")
		for {
			player.Play(dj.NextSong())
		}
	}
}
