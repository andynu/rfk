// Provide information about songs
package library

import (
	"bufio"
	"os"
)

var songsList []string

func init() {
	LoadSongs("./data/mongongo/songs.txt")
}

func LoadSongs(songsTxt string) *[]string {
	f, err := os.Open(songsTxt)
	check(err)
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
