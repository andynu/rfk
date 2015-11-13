// Selects songs for a Player
package dj

import (
	"container/list"
	"errors"
	"math/rand"
	"rfk/library"
	"time"
)

var songs []*library.Song
var requests = list.New()

func NextSong() string {
	if songs == nil {
		songs = *library.Songs()
	}

	next_song, err := requestedSong()
	if err != nil {
		next_song, _ = randomSong()
	}

	return next_song
}

func randomSong() (string, error) {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(songs))
	return songs[idx].Path, nil
}

func requestedSong() (string, error) {
	if requests.Len() >= 0 {
		//return pop(requests).Value.(string), nil
	}
	return "", errors.New("NoRequests")
}

func pop(list *list.List) *list.Element {
	e := list.Back()
	list.Remove(e)
	return e
}
