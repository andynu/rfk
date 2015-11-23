package dj

import (
	"fmt"
	"math/rand"
	"rfk/karma"
	"rfk/library"
	"time"
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
