package dj

import (
	"container/list"
	"errors"
	"log"

	"sync"

	"github.com/andynu/rfk/server/library"
)

type requestList struct {
	list.List
	sync.Mutex
}

var requests requestList = requestList{}

func (list *requestList) popSong() (library.Song, error) {
	if list.Len() > 0 {
		song, ok := requests.pop().Value.(library.Song)
		if ok {
			return song, nil
		}
	}
	return library.Song{}, errors.New("NoRequests")
}

func (list *requestList) pop() *list.Element {
	e := list.Back()
	list.Remove(e)
	return e
}

func remove(song *library.Song) {
	for e := requests.Front(); e != nil; e = e.Next() {
		rsong, _ := e.Value.(library.Song)
		if rsong.Hash == song.Hash {
			requests.Remove(e)
		}
	}
}

func (list *requestList) addAll(songs []*library.Song) {
	for _, song := range songs {
		list.PushBack(*song)
	}
}

func (list *requestList) removeAll(songs []*library.Song) {
	for _, song := range songs {
		remove(song)
	}
}

func PopRequest() (library.Song, error) {
	requests.Lock()
	defer requests.Unlock()

	return requests.popSong()
}

func Request(songs []*library.Song) {
	requests.Lock()
	defer requests.Unlock()

	requests.addAll(songs)
}

func Unrequest(songs []*library.Song) {
	requests.Lock()
	defer requests.Unlock()

	requests.removeAll(songs)
}

func ClearRequests() {
	requests = requestList{}
}

func RequestCount() int {
	requests.Lock()
	defer requests.Unlock()

	return requests.Len()
}

func Requests() []*library.Song {
	var retSongs []*library.Song
	requests.Lock()
	defer requests.Unlock()

	for e := requests.Front(); e != nil; e = e.Next() {
		song, _ := e.Value.(library.Song)
		retSongs = append(retSongs, &song)
	}
	return retSongs
}

// see dj.NextSong()
func requestedSong() (library.Song, error) {
	log.Printf("requestedSong()")
	return PopRequest()
}
