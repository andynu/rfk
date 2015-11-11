// RFK Entry point, and main loop
package main

import (
	"flag"
	"log"
	"rfk/dj"
	"rfk/library"
	"rfk/player"
)

var songsTxt = flag.String("c", "/home/andy/songs.txt", "Songs txt file, one path per line")

func main() {
	flag.Parse()
	library.LoadSongs(*songsTxt)
	log.Println("RFK v3 startup")
	for {
		song := dj.NextSong()
		player.Play(song)
	}
}
