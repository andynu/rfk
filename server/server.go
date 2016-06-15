// RFK Entry point, and main loop
package main

import (
	"flag"
	"log"
	"time"

	"github.com/andynu/rfk/server/api/console"
	"github.com/andynu/rfk/server/api/rest"
	"github.com/andynu/rfk/server/api/rpc"
	"github.com/andynu/rfk/server/config"
	"github.com/andynu/rfk/server/dj"
	"github.com/andynu/rfk/server/env"
	_ "github.com/andynu/rfk/server/env/sensors/weather"
	"github.com/andynu/rfk/server/karma"
	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/observer"
	"github.com/andynu/rfk/server/player"
)

func main() {
	command := flag.String("e", "", "command")
	configPath := flag.String("c", "", "config path")
	dataPath := flag.String("d", "", "data path")
	webPlayerOnly := flag.Bool("webplayer", false, "webplayer; no mpg123 output")
	startPaused := flag.Bool("paused", false, "start paused")
	flag.Parse()

	if *webPlayerOnly {
		player.Silence()
	}

	ensureBinaryExists("mpg123")

	config.Load(*configPath)

	ensureDataPath(*dataPath)
	ensureSongs()

	switch *command {
	case "add":
		library.AddPaths(flag.Args())

	default:

		rpc.RPCListener()
		console.InputListener()
		rest.RESTListener()

		library.Load()
		karma.Setup()
		env.StartEnvUpdater()

		if *startPaused {
			player.TogglePlayPause()
		}

		observer.Observe("player.played", func(msg interface{}) {
			song := msg.(*library.Song)
			log.Printf("Played %v", song)
			karma.Log(song, 1)
		})

		observer.Observe("player.skip", func(msg interface{}) {
			song := msg.(*library.Song)
			log.Printf("Skipped %v", song)
			karma.Log(song, -2)
		})

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
