// Selects songs for a Player
package dj

import (
	"fmt"
	"log"

	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/observer"
)

var Songs library.SongList

var djs = []func() (library.Song, error){
	requestedSong,
	randomNormalSong,
}

var djNames = []string{
	"requests",
	"randomNonNegativeRankSong",
}

func init() {
	observer.Observe("library.loaded", func(msg interface{}) {
		setSongs(library.Songs)
	})
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

func setSongs(songs library.SongList) {
	for _, song := range songs {
		if song.Hash != "" && song.Rank >= 0 {
			Songs = append(Songs, song)
		}
	}
}
