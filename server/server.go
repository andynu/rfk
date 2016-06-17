// RFK Entry point, and main loop
package main

import (
	"flag"

	"github.com/andynu/rfk/server/api/console"
	"github.com/andynu/rfk/server/api/rest"
	"github.com/andynu/rfk/server/api/rpc"
	"github.com/andynu/rfk/server/config"
	"github.com/andynu/rfk/server/dj"
	"github.com/andynu/rfk/server/env"
	_ "github.com/andynu/rfk/server/env/sensors/weather"
	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/player"
)

type Listener interface {
	Listener()
}

func main() {
	command := flag.String("e", "", "command")
	configPath := flag.String("c", "", "config path")
	dataPath := flag.String("d", "", "data path")
	webPlayerOnly := flag.Bool("webplayer", false, "webplayer; no mpg123 output")
	startPaused := flag.Bool("paused", false, "start paused")
	flag.Parse()

	config.Load(*configPath, *dataPath)

	ensureBinaryExists("mpg123")
	ensureSongs()

	switch *command {
	case "add":
		library.AddPaths(flag.Args())

	default:

		library.Load()

		if *webPlayerOnly {
			player.Silence()
		}

		if *startPaused {
			player.Pause()
		}

		nextSongCh := make(chan library.Song, 1)

		go dj.ServeSongs(nextSongCh)
		go player.PlaySongs(nextSongCh)

		go env.Updater()

		go rpc.Listener()
		go console.Listener()
		go rest.Listener()

	} // switch

	doneCh := make(chan struct{})
	<-doneCh
}
