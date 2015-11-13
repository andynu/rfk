// Provide information about songs
package library

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Song struct {
	Hash string
	Path string
}

var songs []Song
var songPathMap map[string]Song
var songHashMap map[string][]Song

func init() {
	host, err := os.Hostname()
	check(err)
	songsPath := fmt.Sprintf("./data/%s/songs.txt", host)
	songHashesPath := fmt.Sprintf("./data/%s/song_hashes.txt", host)
	log.Printf("Loading songs from %q", songsPath)

	songs = make([]Song, 1000)
	songHashMap = make(map[string][]Song, 1000)
	songPathMap = make(map[string]Song, 1000)

	if ok, _ := pathExists(songHashesPath); ok {
		LoadSongHashesMap(songHashesPath)
	} else if ok, _ := pathExists(songsPath); ok {
		LoadSongs(songsPath)
	} else {
		log.Printf("Could not load %q - %q", songHashesPath, err)
	}
}

// consumes a text file of the form "hash\tfilepath\n"
func LoadSongHashesMap(songHashesTxt string) {
	hashIdx := 0
	pathIdx := 1

	f, err := os.Open(songHashesTxt)
	if err != nil {
		log.Print("Could not load song hashes.")
		log.Print(err)
		return
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = '\t'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Print("Could not load song hashes.")
		log.Print(err)
		return
	}

	for _, record := range records {
		var song Song
		song.Path = record[pathIdx]
		song.Hash = record[hashIdx]
		songs = append(songs, song)
		songPathMap[song.Path] = song
		if songHashMap[song.Hash] == nil {
			songHashMap[song.Hash] = make([]Song, 1)
			songHashMap[song.Hash] = append(songHashMap[song.Hash], song)
		}
	}
}

func LoadSongs(songsTxt string) {
	f, err := os.Open(songsTxt)
	if err != nil {
		log.Print("Could not load songs. Please add a list of mp3 paths to the file listed below")
		log.Print(err)
		os.Exit(1)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var song Song
		song.Path = scanner.Text()
		songPathMap[song.Path] = song
		songs = append(songs, song)
	}

}

func Songs() *[]Song {
	return &songs
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
