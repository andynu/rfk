// RFK Entry point, and main loop
package main

import (
	"log"
	"rfk/dj"
	"rfk/player"
)

func main() {
	log.Println("RFK v3 startup")
	for {
		player.Play(dj.NextSong())
	}
}
