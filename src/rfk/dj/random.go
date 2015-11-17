package dj

import (
	"math/rand"
	"rfk/library"
	"time"
)

func randomSong() (library.Song, error) {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(library.Songs) - 1)
	song := *library.Songs[idx]
	return song, nil
}
