// Interprets impressions into library.Song.Rank
package karma

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
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
	hashIdx := 1
	tagIdx := 2
	impressionIdx := 3

	impressionLogPath := path.Join(config.DataPath, "impression.log")
	log.Printf("Reading impressions from: %q", impressionLogPath)
	f, err := os.Open(impressionLogPath)
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
		if len(record) != 4 {
			return fmt.Errorf("Malformed impression.log, wrong number of columns %d expected 4", len(record))
		}
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

	sort.Sort(library.Songs)

	log.Printf("Loading karma: %d impressions loaded", impressionCount)
	return nil
}
