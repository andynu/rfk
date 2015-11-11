package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/dhowden/tag"
)

const concurrency = 100

var sem = make(chan bool, concurrency)
var out = make(chan string)

func main() {
	//var checksum = flag.Bool("m", false, "machine readable (audio only) checksums")
	flag.Parse()

	//runtime.GOMAXPROCS(2)

	// generate audio only checksums from file params or stdin
	files := flag.Args()

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
	check(err)
	checksum, err := tag.Sum(f)
	check(err)
	out <- fmt.Sprintf("%s\t%q", checksum, path)
}

//m, err := tag.ReadFrom(f)
//log.Print(m.Format())
//log.Print(m.Title())
//log.Print(m.Album())
//log.Print(m.Artist())
//log.Print(m.Genre())
//log.Print(m.Year())

func check(err error) {
	if err != nil {
		panic(err)
	}
}
