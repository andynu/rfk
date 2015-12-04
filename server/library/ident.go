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

	"github.com/andynu/rfk/server/config"
	"github.com/dhowden/tag"
)

const concurrency = 10

var sem = make(chan bool, concurrency)
var out = make(chan string)

func IdentifySongs(songs []*Song, outFile string) {

	log.Printf("Identifying songs...")

	count := 0
	if len(songs) != 0 {
		// from Params
		count = len(songs)
		for _, song := range songs {
			if song.Hash != "" {
				count--
				continue
			}
			go audioFileChecksum(song.Path)
		}
	}

	total := count

	songHashesPath := path.Join(config.DataPath, "song_hashes.txt")
	f, err := os.OpenFile(songHashesPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		panic(fmt.Errorf("%q: %v", songHashesPath, err))
	}
	for count != 0 {
		f.WriteString(<-out + "\n")
		count--

		if count%100 == 0 {
			log.Printf("Identified %d songs.\n", total-count)
		}

	}
}

func audioFileChecksum(path string) {
	sem <- true
	defer func() { <-sem }()
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		out <- fmt.Sprintf("%q\t%q", err, path)
		return
	}
	checksum, err := tag.Sum(f)
	if err != nil {
		out <- fmt.Sprintf("%q\t%q", err, path)
		return
	}
	out <- fmt.Sprintf("%s\t%q", checksum, path)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
