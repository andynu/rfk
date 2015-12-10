package library

// generates audio only checksums
//
// INPUT
//
// using the songs.txt (absolute paths to mp3s, one per line)
//
//    /absolute/path/song.mp3
//    /absolute/path/other_song.mp3
//
// e.g.
//
//    find /music -name '*.mp3' > songs.txt
//
// OUTPUT
//
// a tab separated file of the hash, and the absolute path
//
//    8752c89b6212138f21488d4e775123a478a753c2	"/absolute/path/song.mp3"
//    ce370b2059057e7e114206c748bf1b695e928861	"/absolute/path/other_song.mp3"
//
// RUN
//
// To execute the binary:
//
//    cat songs.txt | ./rfk-ident > song_hashes.txt
//
// Individual mp3 hash
//
//    ./rfk-ident /absolute/path/song.mp3
//    8752c89b6212138f21488d4e775123a478a753c2	"/absolute/path/song.mp3"
//

import (
	"fmt"
	"log"
	"os"
	"path"
	"sync"

	"github.com/andynu/rfk/server/config"
	"github.com/dhowden/tag"
)

func IdentifySongs(songs []*Song, outFile string) {

	log.Printf("Identifying songs...")

	// How many identification routines can run concurrently
	const concurrency = 10
	var sem = make(chan bool, concurrency)

	var songsOutCh = make(chan *Song)
	var wg sync.WaitGroup

	i := 0
	if len(songs) != 0 {
		// from Params
		for _, song := range songs {
			if song.Hash == "" {
				i++
				wg.Add(1)
				go func(song *Song) {
					sem <- true
					defer func() { <-sem }()
					defer func() { wg.Done() }()
					hash, err := audioFileChecksum(song.Path)
					if err != nil {
						log.Printf("ident: err: %v for %s", err, song.Path)
						return
					}
					song.Hash = hash
					songsOutCh <- song
				}(song)

			}
		}
	}

	// Close the output channel when all the audioFileCheksums complete.
	go func() {
		wg.Wait()
		close(songsOutCh)
	}()

	go func() {
		songHashesPath := path.Join(config.DataPath, "song_hashes.txt")
		f, err := os.OpenFile(songHashesPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
		if err != nil {
			panic(fmt.Errorf("%q: %v", songHashesPath, err))
		}

		i = 0
		for song := range songsOutCh {
			f.WriteString(fmt.Sprintf("%s\t%q\n", song.Hash, song.Path))
			i++
			if i%100 == 0 {
				log.Printf("Attempted identification of %d songs\n", i)
			}
		}
		log.Printf("Identifying songs...done\n")
	}()
}

func audioFileChecksum(path string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	checksum, err := tag.Sum(f)
	if err != nil {
		return "", err
	}
	return checksum, nil
}

// if something goes wrong, you get a blank one.
func metaData(path string) SongMeta {
	var meta SongMeta

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return meta
	}

	m, err := tag.ReadFrom(f)
	if err != nil {
		return meta
	}

	meta.Title = m.Title()
	meta.Album = m.Album()
	meta.Artist = m.Artist()

	return meta
}
