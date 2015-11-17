package library

import (
	"bufio"
	"log"
	"os"
)

func loadSongs(songsTxt string) error {
	log.Printf("Loading songs from %q", songsTxt)
	f, err := os.Open(songsTxt)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		song := Song{Path: scanner.Text()}
		songPathMap[song.Path] = &song
		Songs = append(Songs, &song)
	}
	return nil
}
