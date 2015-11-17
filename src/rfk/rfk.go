// RFK Entry point, and main loop
package main

import (
	"flag"
	"log"
	"rfk/dj"
	"rfk/karma"
	"rfk/library"
	"rfk/observer"
	"rfk/player"
	"rfk/rpc"
	"time"
)

func main() {
	command := flag.String("c", "", "command")
	flag.Parse()

	switch *command {
	case "karma":
		karma.LoadImpressions()
	case "add":
		library.AddPaths(flag.Args())
	default:

		rpc.SetupRPC()
		listenForInput()
		library.Load()

		observer.Observe("player.played", func(msg interface{}) {
			song := msg.(library.Song)
			log.Printf("Played %v", song)
			karma.Log(song, 1)
		})

		observer.Observe("player.skip", func(msg interface{}) {
			song := msg.(library.Song)
			log.Printf("Played %v", song)
			karma.Log(song, -2)
		})

		for {

			song, err := dj.NextSong()
			if err != nil {
				log.Printf("rfk: %v", err)
				time.Sleep(1 * time.Second)
				continue
			}

			err = player.Play(song)
			if err != nil {
				log.Printf("rfk: %v", err)
				time.Sleep(1 * time.Second)
			}
		}
	}
}
