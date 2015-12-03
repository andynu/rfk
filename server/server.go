// RFK Entry point, and main loop
package main

import (
	"flag"
	"log"
	"time"

	"github.com/andynu/rfk/server/config"
	"github.com/andynu/rfk/server/console"
	"github.com/andynu/rfk/server/dj"
	"github.com/andynu/rfk/server/env"
	_ "github.com/andynu/rfk/server/env/sensors/weather"
	"github.com/andynu/rfk/server/karma"
	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/observer"
	"github.com/andynu/rfk/server/player"
	"github.com/andynu/rfk/server/rpc"
)

func main() {
	command := flag.String("e", "", "command")
	configPath := flag.String("c", "", "config path")
	flag.Parse()

	config.Load(configPath)

	switch *command {
	case "add":
		library.AddPaths(flag.Args())

	default:

		checkPrereqs()
		console.InputListener()
		rpc.SetupRPC()

		library.Load()
		karma.Load()

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

		env.Prime()
		env.LogFull()

		for {

			if !player.IsPlaying() {
				time.Sleep(1 * time.Second)
				continue
			}

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
			//panic("howmanygoroutines?")
		}
	}
}
