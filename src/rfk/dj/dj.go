// Selects songs for a Player
package dj

import (
	"math/rand"
	"rfk/library"
	"time"
)

func NextSong() string {
	return randomSong()
}

func randomSong() string {
	rand.Seed(time.Now().UnixNano())
	songs := library.Songs()
	idx := rand.Intn(len(songs))
	return songs[idx]
}
