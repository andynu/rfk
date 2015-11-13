// Provide information about songs
package library

import (
	"bufio"
	"encoding/csv"
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

// consumes a text file of the form "hash\tfilepath\n"
func LoadSongIdMap(songHashesTxt string) map[string]string {
	var pathIdMap = make(map[string]string, 1000)
	hashIdx := 0
	pathIdx := 1

	f, err := os.Open(songHashesTxt)
	if err != nil {
		log.Print("Could not load song hashes.")
		log.Print(err)
		return map[string]string(nil)
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = '\t'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Print("Could not load song hashes.")
		log.Print(err)
		return map[string]string(nil)
	}

	for _, record := range records {
		pathIdMap[record[pathIdx]] = record[hashIdx]
	}

	fmt.Println(pathIdMap)

	return pathIdMap
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
