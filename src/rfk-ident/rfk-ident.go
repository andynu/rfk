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
		for _, path := range files {
			go audioFileChecksum(path)
		}
	} else {
		// from Stdin
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			path := input.Text()
			//log.Printf("scan %d. %q", count, path)
			go audioFileChecksum(path)
			count++
		}
	}

	for count != 0 {
		count--
		fmt.Println(<-out)
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
