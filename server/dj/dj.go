// Selects songs for a Player
package dj

import (
	"fmt"
	"log"
	"time"

	"github.com/andynu/rfk/server/dj/listened"
	"github.com/andynu/rfk/server/library"
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
	//observer.Observe("karma.loaded", func(msg interface{}) {
	//	log.Printf("dj songs set")
	//})
}

func ServeSongs(nextSongCh chan library.Song) {
	for {
		song, err := NextSong()
		if err != nil {
			log.Printf("rfk: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		nextSongCh <- song
	}
}

func NextSong() (library.Song, error) {

	var next_song library.Song
	var err error

	for i, dj := range djs {
		next_song, err = dj()

		if listened.Includes(next_song) {
			log.Printf("Skipping repeat song: %v", next_song)
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

func SetSongs(songs library.SongList) {
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
