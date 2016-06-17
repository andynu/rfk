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
	"time"

	"github.com/andynu/rfk/server/config"
	"github.com/andynu/rfk/server/historicalEnv"
	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/observer"
)

var mu sync.Mutex
var logger *log.Logger

// Sum(impressions) by Song.Hash
var SongKarma map[string]int

var graph *library.Graph

func init() {
	observer.Observe("config.loaded", func(msg interface{}) {
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
	})

	observer.Observe("player.played", func(msg interface{}) {
		song := msg.(*library.Song)
		log.Printf("Played %v", song)
		Log(song, 1)
	})

	observer.Observe("player.skip", func(msg interface{}) {
		song := msg.(*library.Song)
		log.Printf("Skipped %v", song)
		Log(song, -2)
	})

}

// Record a positive/negative impression of a Song.
func Log(song *library.Song, impression int) {
	song.Rank += float64(impression)
	logImpression(song, "karma", impression)
}

func logImpression(song *library.Song, tag string, impression int) {
	if song.Hash == "" {
		log.Printf("Cannot take impression for %q, no song hash.\n", song.Path)
		return
	}
	mu.Lock()
	logger.Printf("\t%s\t%s\t%d", "karma", song.Hash, impression)
	mu.Unlock()
}

func LogTag(song *library.Song, tag string) {
	mu.Lock()
	logger.Printf("\t%s\t%s\t%s", "tag", song.Hash, tag)
	mu.Unlock()
}

func LogEnv(sensor string, value string) {
	mu.Lock()
	logger.Printf("\t%s\t%s\t%s", "env", sensor, value)
	mu.Unlock()
}

// Build up the SongKarma map, and spread impressiosn to library.Song.Rank
func Load() error {
	SongKarma = make(map[string]int)

	impressionCount := 0

	timestampIdx := 0
	logTypeIdx := 1

	// logType = karma
	hashIdx := 2
	impressionIdx := 3

	// logType = env
	tagIdx := 2
	valIdx := 3

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
		timestampStr := record[timestampIdx]
		timestamp, _ := time.Parse("2016/02/21 20:18:49.091737", timestampStr)
		logType := record[logTypeIdx]
		switch logType {
		case "karma":
			hash := record[hashIdx]
			impression, _ := strconv.Atoi(record[impressionIdx])
			SongKarma[hash] += impression

			song, err := library.ByHash(hash)

			if err != nil {
				continue
			}

			impressionCount++
			historicalEnv.Impress(timestamp, song, impression)
			go song.PathGraphImpress(impression)

		case "env":
			sensor := record[tagIdx]
			val := record[valIdx]
			historicalEnv.Update(timestamp, sensor, val)
		}
	}
	//historicalEnv.Print()
	//log.Printf("%v", SongKarma)

	sort.Sort(library.Songs)

	observer.Notify("karma.loaded", struct{}{})
	log.Printf("Loading karma: %d impressions loaded", impressionCount)
	return nil
}
