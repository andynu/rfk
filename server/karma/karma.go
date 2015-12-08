// Interprets impressions into library.Song.Rank
package karma

import (
	"encoding/csv"
	"log"
	"os"
	"path"
	"strconv"
	"sync"

	"github.com/andynu/rfk/server/config"
	"github.com/andynu/rfk/server/library"
)

var mu sync.Mutex
var logger *log.Logger

// Sum(impressions) by Song.Hash
var SongKarma map[string]int

var graph *library.Graph

func Setup() {
	logfile, err := os.OpenFile(path.Join(config.DataPath, "impression.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if err != nil {
		panic(err)
	}
	logger = log.New(logfile, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	err = Load()
	if err != nil {
		panic(err)
	}

}

// Record a positive/negative impression of a Song.
func Log(song library.Song, impression int) {
	LogTag(song, "karma", impression)
}

func LogTag(song library.Song, tag string, impression int) {
	if song.Hash == "" {
		log.Printf("Cannot take impression for %q, no song hash.\n", song.Path)
		return
	}
	mu.Lock()
	logger.Printf("\t%s\t%s\t%d", song.Hash, "karma", impression)
	mu.Unlock()
}

// Build up the SongKarma map, and spread impressiosn to library.Song.Rank
func Load() error {
	SongKarma = make(map[string]int)

	impressionCount := 0

	timestampIdx := 0
	hashIdx := 0
	tagIdx := 1
	impressionIdx := 2

	f, err := os.Open(path.Join(config.DataPath, "impression.log"))
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
		_ = record[timestampIdx]
		hash := record[hashIdx]
		_ = record[tagIdx]
		impression, _ := strconv.Atoi(record[impressionIdx])
		SongKarma[hash] += impression

		song, err := library.ByHash(hash)
		if err == nil {
			//log.Printf("spread it %v %v", hash, impression)
			impressionCount++
			go song.PathGraphImpress(impression)
		}
	}
	//log.Printf("%v", SongKarma)
	log.Printf("Loading karma: %d impressions loaded", impressionCount)
	return nil
}
