package api

import (
	"github.com/andynu/rfk/server/dj"
	"github.com/andynu/rfk/server/karma"
	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/player"
)

type PlayerStatusResult struct {
	CurrentSong     library.Song
	CurrentSongMeta library.SongMeta
	LastSong        library.Song
	LastSongMeta    library.SongMeta
	PlayPauseState  string
	RequestCount    int
}

func PlayerStatus() PlayerStatusResult {
	var result PlayerStatusResult
	result.CurrentSong = player.CurrentSong
	result.CurrentSongMeta = player.CurrentSong.Meta()
	result.LastSong = player.LastSong
	result.LastSongMeta = player.LastSong.Meta()
	result.PlayPauseState = player.PlayPauseState()
	result.RequestCount = dj.RequestCount()
	return result
}

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
	karma.LogTag(player.CurrentSong, tag)
}

func TagLastSong(tag string) {
	karma.LogTag(player.LastSong, tag)
}
