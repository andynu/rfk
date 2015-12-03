package dj

import (
	"container/list"
	"errors"

	"github.com/andynu/rfk/server/library"

	"log"
)

var requests *list.List
func init() {
	requests = list.New()
}

func requestedSong() (library.Song, error) {
	if requests.Len() > 0 {
		log.Printf("requestCount: %d", requests.Len())
		val := pop(requests).Value
		song, ok := val.(library.Song)
		if ok {
			return song, nil
		}
	}
	return library.Song{}, errors.New("NoRequests")
}

func pop(list *list.List) *list.Element {
	e := list.Back()
	list.Remove(e)
	return e
}

func Request(songs []*library.Song) {
	for _, song := range songs {
		requests.PushBack(*song)
	}
}
