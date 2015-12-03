package dj

import (
	"container/list"
	"errors"

	"github.com/andynu/rfk/server/library"
	"sync"
)

var requests_mu sync.Mutex
var requests = list.New()

func requestedSong() (library.Song, error) {
	requests_mu.Lock()
	defer func() { requests_mu.Unlock() }()
	if requests.Len() >= 0 {
		song := pop(requests).Value.(library.Song)
		return song, nil
	}
	return library.Song{}, errors.New("NoRequests")
}

func pop(list *list.List) *list.Element {
	e := list.Back()
	list.Remove(e)
	return e
}

func Request(songs []*library.Song) {
	requests_mu.Lock()
	for _, song := range songs {
		requests.PushBack(*song)
	}
	requests_mu.Unlock()
}

func ClearRequests() {
	requests_mu.Lock()
	requests = list.New()
	requests_mu.Unlock()
}
