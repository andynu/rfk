// RFK Entry point, and main loop
package main

import (
	"log"
	"rfk/dj"
	"rfk/player"
	"flag"
)

var songsTxt = flag.String("c", "/home/andy/songs.txt", "Songs txt file, one path per line")

func main() {
  flag.Parse()
	log.Println("RFK v3 startup")
	for {
		song := dj.NextSong(*songsTxt)
		player.Play(song)
	}
}
