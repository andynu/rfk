package karma

import (
	"encoding/csv"
	"log"
	"os"
	"path"
	"rfk/config"
	"strconv"
)

var SongKarma map[string]int

func LoadImpressions() error {
	SongKarma = make(map[string]int)

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
	}
	//log.Printf("%v", SongKarma)
	log.Printf("karma: impressions loaded")
	return nil
}