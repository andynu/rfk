// Selects songs for a Player
package dj

import (
	"fmt"
	"log"

	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/observer"
	"github.com/andynu/rfk/server/dj/listened"
)

var Songs library.SongList

var djs = []func() (library.Song, error){
	requestedSong,
	randomNormalSong,
}

var djNames = []string{
	"requests",
	"randomNormalSong",
}

func init() {
	observer.Observe("karma.loaded", func(msg interface{}) {
		setSongs(library.Songs)
	})
}

func NextSong() (library.Song, error) {

	var next_song library.Song
	var err error

	for i, dj := range djs {
		next_song, err = dj()

		if listened.Includes(next_song) {
			continue
		}

		if err == nil {
			listened.Add(next_song)
			log.Printf("")
			log.Printf("Using DJ:%v:%v", i, djNames[i])
			return next_song, nil
		}
	}




	return library.Song{}, fmt.Errorf("DJFail")
}

func setSongs(songs library.SongList) {
	considered := 0
	selected := 0
	for _, song := range songs {
		considered++
		if song.Hash != "" && song.Rank >= 0 {
			selected++
			Songs = append(Songs, song)
		}
	}
	fmt.Printf("dj songs: considered=%d selected=%d\n", considered, selected)
}
