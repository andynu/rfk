package dj

import (
	"math/rand"
	"time"

	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/observer"
)

var Songs library.SongList

func init() {
	observer.Observe("library.loaded", func(msg interface{}) {
		setSongs(library.Songs)
	})
}

func setSongs(songs library.SongList) {
	for _, song := range songs {
		if song.Hash != "" && song.Rank >= 0 {
			Songs = append(Songs, song)
		}
	}
}

func randomSong() (library.Song, error) {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(Songs) - 1)
	song := *Songs[idx]
	return song, nil
}

func randomNormalSong() (library.Song, error) {
	rand.Seed(time.Now().UnixNano())
	idx := normalRand(len(Songs) - 1)
	song := *Songs[idx]
	return song, nil
}

// a random normal as sampled from a normal distribution.
// bound between [0, max]
func normalRand(max int) int {
	maxf := float64(max)
	mean := float64(maxf / 2.0)
	stdDev := float64(maxf / 2.0)
	sample := rand.NormFloat64()*stdDev + mean
	if sample < 0 {
		sample = 0.0
	}
	if sample > maxf {
		sample = maxf
	}
	return int(sample)
}
