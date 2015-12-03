package dj

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/andynu/rfk/server/karma"
	"github.com/andynu/rfk/server/library"
)

func randomSong() (library.Song, error) {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(library.Songs) - 1)
	song := *library.Songs[idx]
	return song, nil
}

// karma is song specific
func randomNonNegativeKarmaSong() (library.Song, error) {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(library.Songs) - 1)
	song := *library.Songs[idx]
	if karma.SongKarma[song.Hash] < 0 {
		return library.Song{}, fmt.Errorf("NegSong")
	}
	return song, nil
}

// rank incorporates the karma of nearbye songs
func randomNonNegativeRankSong() (library.Song, error) {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(library.Songs) - 1)
	song := *library.Songs[idx]
	if song.Rank < 0 {
		return library.Song{}, fmt.Errorf("NegSong")
	}
	return song, nil
}

func randomNormalNonNegRankSong() (library.Song, error) {
	rand.Seed(time.Now().UnixNano())
	idx := normalRand(len(library.Songs) - 1)
	song := *library.Songs[idx]
	if song.Rank < 0 {
		return library.Song{}, fmt.Errorf("NegSong")
	}
	return song, nil
}

func randomNormalRankSong() (library.Song, error) {
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
