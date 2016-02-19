package library

import (
	"bufio"
	"log"
	"os"
)

func loadSongErrors(songsErrorsTxtPath string) error {
	log.Printf("Loading songs error paths from %q", songsErrorsTxtPath)
	f, err := os.Open(songsErrorsTxtPath)
	if err != nil {
		return err
	}
	defer f.Close()

	var paths []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		path := scanner.Text()
		paths = append(paths, path)
	}

	SongErrorPaths := make(map[string]bool, len(paths))
	for _, path := range paths {
		SongErrorPaths[path] = true
	}

	log.Printf("Loaded %d error paths", len(SongErrorPaths))

	return nil
}

func loadSongs(songsTxt string) error {
	log.Printf("Loading songs from %q", songsTxt)
	f, err := os.Open(songsTxt)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		path := scanner.Text()
		if SongErrorPaths[path] {
			log.Printf("Skipping error path: %q", path)
		} else {
			song := Song{Path: path}
			songPathMap[song.Path] = &song
			Songs = append(Songs, &song)
		}
	}
	log.Printf("Loaded %d songs (paths: %d)", len(Songs), len(songPathMap))
	return nil
}
