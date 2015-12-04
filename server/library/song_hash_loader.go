package library

import (
	"encoding/csv"
	"log"
	"os"
)

// consumes a text file of the form "hash\tfilepath\n"
// the same format that rfk-ident produces
func loadSongHashesMap(songHashesTxt string) error {
	log.Printf("Loading songs from %q", songHashesTxt)
	hashIdx := 0
	pathIdx := 1

	f, err := os.Open(songHashesTxt)
	if err != nil {
		return err
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = '\t'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {
    song := songPathMap[record[pathIdx]]

    if song == nil {
      song = &Song{Path: record[pathIdx], Hash: record[hashIdx]}
      Songs = append(Songs, song)
      songPathMap[song.Path] = song
    }

    song.Hash = record[hashIdx]
    songHashMap[song.Hash] = append(songHashMap[song.Hash], song)
	}
	return nil
}
