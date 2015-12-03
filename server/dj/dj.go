// Selects songs for a Player
package dj

import (
	"fmt"
	"log"

	"github.com/andynu/rfk/server/library"
)

var djs = []func() (library.Song, error){
	requestedSong,
	randomNonNegativeRankSong,
	//karmaSong(),
}

var djNames = []string{
	"requests",
	"randomNonNegativeRankSong",
}

func NextSong() (library.Song, error) {

	var next_song library.Song
	var err error

	for i, dj := range djs {
		next_song, err = dj()
		if err == nil {
			log.Printf("")
			log.Printf("Using DJ:%v:%v", i, djNames[i])
			return next_song, nil
		}
	}

	return library.Song{}, fmt.Errorf("DJFail")
}
