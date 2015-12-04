// Provide information about songs
package library

import (
	"fmt"
	"log"
	"path"

	"github.com/andynu/rfk/server/config"
	"github.com/andynu/rfk/server/observer"
	"strings"
)

// The list of Songs
var Songs []*Song

// Map from path to Song
var songPathMap map[string]*Song

// Map from hash to Song
var songHashMap map[string][]*Song

var graph *Graph

// The root Nodes for the path graph
var PathRoots []Node

// song file paths, used by Load() and AddPaths()
var songsPath, songHashesPath string

// loads Songs from either song_hashes.txt or songs.txt (first to exist).
// loads the path graph
func Load() {

	songHashMap = make(map[string][]*Song, 1000)
	songPathMap = make(map[string]*Song, 1000)

	if Songs == nil {
		songHashesPath = path.Join(config.DataPath, "song_hashes.txt")
		err := loadSongHashesMap(songHashesPath)
		panicOnErr(err)
	}

	if Songs == nil {
		songsPath = path.Join(config.DataPath, "songs.txt")
		err := loadSongs(songsPath)
		panicOnErr(err)
	}

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

func panicOnErr(err error) {
	if err != nil {
		panic(fmt.Errorf("library: %v", err))
	}
}
