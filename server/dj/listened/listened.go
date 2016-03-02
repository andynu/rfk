package listened

import (
	"github.com/andynu/rfk/server/library"
)

const buffSize = 200
var listenedSongs [buffSize]library.Song
var idx int

func Add(song library.Song){
	listenedSongs[idx] = song

	idx++
	if idx >= buffSize {
		idx = 0
	}
}

func Includes(song library.Song) bool {
	for _, listenedSong := range listenedSongs {
		if song.Hash == listenedSong.Hash {
			return true
		}
	}
	return false
}


