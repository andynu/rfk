// Provide information about songs
package library

import (
	"fmt"
	"log"
	"path"
	"rfk/config"
	"rfk/observer"
)

type Song struct {
	Hash string
	Path string
}

var songsPath, songHashesPath string

var Songs []*Song
var songPathMap map[string]*Song
var songHashMap map[string][]*Song

func Load() {

	songHashMap = make(map[string][]*Song, 1000)
	songPathMap = make(map[string]*Song, 1000)

	if Songs == nil {
		songHashesPath = path.Join(config.Config.DataPath, "song_hashes.txt")
		err := loadSongHashesMap(songHashesPath)
		panicOnErr(err)
	}

	//if Songs == nil {
	//	songsPath = path.Join(config.Config.DataPath, "songs.txt")
	//	err := loadSongs(songsPath)
	//	panicOnErr(err)
	//}

	observer.Notify("library.loaded", Song{})
	log.Printf("Loaded %d songs", len(Songs))
}

func ByHash(hash string) (*Song, error) {
	songs := songHashMap[hash]
	log.Printf("DEBUG: songs:%v", songs)
	if songs != nil {
		return songs[0], nil
	}
	return &Song{}, fmt.Errorf("UnknownHash")
}

func panicOnErr(err error) {
	if err != nil {
		panic(fmt.Errorf("library: %v", err))
	}
}
