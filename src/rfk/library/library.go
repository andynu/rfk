// Provide information about songs
package library

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var songsList []string

func init() {
	host, err := os.Hostname()
	check(err)
	songsPath := fmt.Sprintf("./data/%s/songs.txt", host)
	log.Printf("Loading songs from %q", songsPath)
	LoadSongs(songsPath)
}

func LoadSongs(songsTxt string) *[]string {
	f, err := os.Open(songsTxt)
	if err != nil {
		log.Print("Could not load songs. Please add a list of mp3 paths to the file listed below")
		log.Print(err)
		os.Exit(1)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		songsList = append(songsList, scanner.Text())
	}

	return &songsList
}

func Songs() *[]string {
	return &songsList
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
