// Selects songs for a Player
package dj

import (
	"math/rand"
	"rfk/library"
	"time"
)

var songs []string

func NextSong() string {
	if songs == nil {
		songs = *library.Songs()
	}
	return randomSong()
}

func randomSong() string {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(songs))
	return songs[idx]
}
