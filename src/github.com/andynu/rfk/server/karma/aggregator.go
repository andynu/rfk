package karma

import (
	"encoding/csv"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/andynu/rfk/server/config"
	"github.com/andynu/rfk/server/library"
)

// Sum(impressions) by Song.Hash
var SongKarma map[string]int

var graph *library.Graph

// Build up the SongKarma map, and spread impressiosn to library.Song.Rank
func Load() error {
	SongKarma = make(map[string]int)

	impressionCount := 0
	hashIdx := 0
	impressionIdx := 1

	f, err := os.Open(path.Join(config.Config.DataPath, "impression.log"))
	if err != nil {
		panic(err)
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
