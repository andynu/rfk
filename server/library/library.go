// Provide information about songs
package library

import (
	"fmt"
	"log"
	"path"

	"strings"

	"github.com/andynu/rfk/server/config"
	"github.com/andynu/rfk/server/observer"
)

// The list of Songs
type SongList []*Song

func (slice SongList) Len() int {
	return len(slice)
}

func (slice SongList) Less(i, j int) bool {
	return slice[i].Rank < slice[j].Rank
}

func (slice SongList) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (songs SongList) Add(song *Song) {
	songs = append(songs, song)
}

func (songs SongList) Filter(filter func(*Song) bool) *SongList {
	var filteredSongs SongList
	for _, song := range songs {
		if filter(song) {
			filteredSongs.Add(song)
		}
	}
	return &filteredSongs
}

type SongErrorList map[string]bool

var Songs SongList

var SongErrorPaths SongErrorList

// Map from path to Song
var songPathMap map[string]*Song

// Map from hash to Song
var songHashMap map[string][]*Song

var graph *Graph

// The root Nodes for the path graph
var PathRoots []Node

// loads Songs from either song_hashes.txt or songs.txt (first to exist).
// loads the path graph
func Load() {

	songHashMap = make(map[string][]*Song, 1000)
	songPathMap = make(map[string]*Song, 1000)

	songHashesPath := path.Join(config.DataPath, "song_hashes.txt")
	//songsPath := path.Join(config.DataPath, "songs.txt")
	songErrorsPath := path.Join(config.DataPath, "song_errors.txt")

	loadSongErrors(songErrorsPath)
	//loadSongs(songsPath)
	loadSongHashesMap(songHashesPath)
	go IdentifySongs(Songs, songHashesPath)

	observer.Notify("library.loaded", struct{}{})
	log.Printf("Loaded %d songs", len(Songs))
	PathRoots = graph.LoadGraph(Songs)

}

// lookup Song by Song.Hash string
func ByHash(hash string) (*Song, error) {
	songs := songHashMap[hash]
	if songs != nil {
		return songs[0], nil
	}
	return &Song{}, fmt.Errorf("UnknownHash")
}

// lookup Song by Song.Hash string
func ByHashes(hashes []string) ([]*Song, error) {
	var songs []*Song
	for _, hash := range hashes {
		hashSongs := songHashMap[hash]
		if hashSongs != nil {
			songs = append(songs, hashSongs[0])
		}
	}
	return songs, nil
}

// find songs by path substring
func Search(term string) []*Song {
	var songs []*Song
	for _, song := range Songs {
		if strings.Contains(song.Path, term) {
			songs = append(songs, song)
		}
	}
	return songs
}

func Artists() []string {
	artist_map := make(map[string]bool, 1000)
	for _, song := range Songs {
		//fmt.Printf("artists: %v\n", song)
		//fmt.Printf("artists: %v\n", song.Artist)
		artist_map[song.Artist] = true
	}

	var artists []string
	for artist := range artist_map {
		artists = append(artists, artist)
	}
	return artists
}

func panicOnErr(err error) {
	if err != nil {
		panic(fmt.Errorf("library: %v", err))
	}
}
