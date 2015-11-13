// RFK Entry point, and main loop
package main

import (
	"flag"
	"log"
	"rfk/dj"
	"rfk/library"
	"rfk/player"
)

func main() {
	command := flag.String("c", "", "command")
	flag.Parse()

	switch *command {
	case "hashes":
		library.LoadSongIdMap("./data/mongongo/song_hashes.txt")
	default:
		log.Println("RFK v3 startup")
		for {
			player.Play(dj.NextSong())
		}
	}
}
