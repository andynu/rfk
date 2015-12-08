package dj

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/andynu/rfk/server/library"
)

func noHashFilter(djFunc func() (library.Song, error)) func() (library.Song, error) {
	return func() (library.Song, error) {
		song, err := djFunc()
		if err != nil {
			return library.Song{}, err
		}
		if song.Hash == "" {
			return library.Song{}, fmt.Errorf("NoHashSong")
		}
		return song, nil
	}
}

func noNegFilter(djFunc func() (library.Song, error)) func() (library.Song, error) {
	return func() (library.Song, error) {
		song, err := djFunc()
		if err != nil {
			return library.Song{}, err
		}
		if song.Rank < 0 {
			return library.Song{}, fmt.Errorf("NegSong")
		}
		return song, nil
	}
}

func randomSong() (library.Song, error) {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(library.Songs) - 1)
	song := *library.Songs[idx]
	return song, nil
}

func randomNormalSong() (library.Song, error) {
	rand.Seed(time.Now().UnixNano())
	idx := normalRand(len(library.Songs) - 1)
	song := *library.Songs[idx]
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
