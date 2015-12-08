package api

import (
	"github.com/andynu/rfk/server/dj"
	"github.com/andynu/rfk/server/karma"
	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/player"
)

func PlayPause() {
	player.TogglePlayPause()
}

func SkipNoPunish() {
	player.Stop()
}

func Skip() {
	player.Skip()
}

func Reward() {
	karma.Log(player.CurrentSong, 1)
}

func Search(term string) []*library.Song {
	return library.Search(term)
}

func SearchRequest(term string) []*library.Song {
	songs := library.Search(term)
	dj.Request(songs)
	return songs
}

func ClearRequests() {
	dj.ClearRequests()
}

func TagCurrentSong(tag string) {
}

func TagLastSong(tag string) {
}
