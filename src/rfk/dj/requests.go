package dj

import (
	"container/list"
	"errors"
	"rfk/library"
)

var requests = list.New()

func requestedSong() (library.Song, error) {
	if requests.Len() >= 0 {
		//return pop(requests).Value.(string), nil
	}
	return library.Song{}, errors.New("NoRequests")
}

func pop(list *list.List) *list.Element {
	e := list.Back()
	list.Remove(e)
	return e
}
