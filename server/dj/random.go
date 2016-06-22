package dj

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/andynu/rfk/server/library"
)

func randomSong() (library.Song, error) {
	log.Printf("randomSong()")
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(Songs) - 1)
	song := *Songs[idx]
	return song, nil
}

func randomNormalSong() (library.Song, error) {
	log.Printf("randomNormalSong()")
	rand.Seed(time.Now().UnixNano())
	lensongs := len(Songs)
	if lensongs == 0 {
		return library.Song{}, fmt.Errorf("No songs")
	}
	idx := normalRand(lensongs - 1)
	fmt.Printf("randomNormalSong: idx:%v len:%df\n", idx, len(Songs))
	song := *Songs[idx]
	fmt.Printf("randomNormalSong: %v %f\n", song, song.Rank)
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
