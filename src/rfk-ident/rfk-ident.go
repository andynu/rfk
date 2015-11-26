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
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dhowden/tag"
)

const concurrency = 10

var sem = make(chan bool, concurrency)
var out = make(chan string)

func main() {

	// generate audio only checksums from file params or stdin
	files := os.Args[1:]

	count := 0
	if len(files) != 0 {
		// from Params
		count = len(files)
		for _, path := range files {
			go audioFileChecksum(path)
		}
	} else {
		// from Stdin
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			path := input.Text()
			go audioFileChecksum(path)
			count++
		}
	}

	for count != 0 {
		fmt.Println(<-out)
		count--
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
