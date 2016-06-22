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
	artistIdx := 2
	albumIdx := 3
	titleIdx := 4

	f, err := os.Open(songHashesTxt)
	if err != nil {
		return err
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comment = '#'
	r.LazyQuotes = true

	records, err := r.ReadAll()
	if err != nil {
		log.Printf("ERROR: %v", err)
		panic(err)
	}

	for _, record := range records {
		path := record[pathIdx]
		song := songPathMap[path]

		if SongErrorPaths[path] {
			log.Printf("Skipping error path: %q", path)
		} else {

			if song == nil {
				song = &Song{
					Path: path,
					Hash: record[hashIdx],
					SongMeta: SongMeta{
						Artist: record[artistIdx],
						Album:  record[albumIdx],
						Title:  record[titleIdx],
					},
				}
				Songs = append(Songs, song)
				songPathMap[song.Path] = song
			}

			song.Hash = record[hashIdx]
			songHashMap[song.Hash] = append(songHashMap[song.Hash], song)
		}
	}
	log.Printf("Loaded %d songs (paths: %d)", len(Songs), len(songPathMap))
	return nil
}
