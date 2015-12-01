// RFK Entry point, and main loop
package main

import (
	"flag"
	"fmt"
	"github.com/andynu/rfk/server/config"
	"github.com/andynu/rfk/server/dj"
	"github.com/andynu/rfk/server/karma"
	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/observer"
	"github.com/andynu/rfk/server/player"
	"github.com/andynu/rfk/server/rpc"
	"log"
	"time"
)

func main() {
	command := flag.String("e", "", "command")
	configPath := flag.String("c", "", "config path")
	flag.Parse()

	config.Load(configPath)

	switch *command {
	case "graph":
		library.Load()
		karma.Load()
		for root := range library.PathRoots {
			fmt.Printf("root : %v\n", root)
			fmt.Printf("song: %v\n", library.Songs[0])
			fmt.Printf("song: %v\n", &library.Songs[0])
			//library.CountImpressionsByDepth(library.Songs[0])
			//library.SpreadImpressionByPath(library.Songs[0], 1)
			for _, song := range library.Songs {
				if song.Rank == 0.0 {
					continue
				}
				fmt.Printf("%f\t%s\n", song.Rank, song.Path)
			}
		}
	case "karma":
		karma.Load()
	case "add":
		library.AddPaths(flag.Args())
	default:

		rpc.SetupRPC()
		listenForInput()
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
			//panic("howmanygoroutines?")
		}
	}
}
