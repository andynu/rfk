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
	logger = log.New(logfile, "", 0)

	err = Load()
	if err != nil {
		panic(err)
	}

}

// Record a positive/negative impression of a Song.
func Log(song library.Song, impression int) {
	if song.Hash == "" {
		log.Printf("Cannot take impression for %q, no song hash.\n", song.Path)
		return
	}
	mu.Lock()
	logger.Printf("%s\t%d", song.Hash, impression)
	mu.Unlock()
}

// Build up the SongKarma map, and spread impressiosn to library.Song.Rank
func Load() error {
	SongKarma = make(map[string]int)

	impressionCount := 0
	hashIdx := 0
	impressionIdx := 1

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
		hash := record[hashIdx]
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
