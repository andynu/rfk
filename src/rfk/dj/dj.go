// Selects songs for a Player
package dj

import (
	"fmt"
	"log"
	"rfk/library"
)

var djs = []func() (library.Song, error){
	requestedSong,
	//karmaSong(),
	randomNonNegativeRankSong,
}

func NextSong() (library.Song, error) {

	var next_song library.Song
	var err error

	for i, dj := range djs {
		next_song, err = dj()
		if err == nil {
			log.Printf("Using DJ:%v", i)
			return next_song, nil
		}
	}

	return library.Song{}, fmt.Errorf("DJFail")
}
