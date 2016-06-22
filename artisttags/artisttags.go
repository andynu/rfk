package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/andynu/rfk/server/config"
	"github.com/andynu/rfk/server/library"
	"github.com/shkh/lastfm-go/lastfm"
)

type lastfmConfig struct {
	Key    string
	Secret string
}

type lookupApi struct {
	api *lastfm.Api
}

func main() {

	library.Load()

	api := lookupApi{}
	api.ConnectWithConfig("lastfm.json")

	artists := library.Artists()
	artist_count := len(artists)
	fmt.Printf("Artist Count: %d\n", artist_count)

	outPath := path.Join(config.DataPath, "artist_tags.txt")
	f, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		panic(fmt.Errorf("%q: %v", outPath, err))
	}
	csv := csv.NewWriter(f)

	ticker := time.NewTicker(time.Millisecond * 201)
	i := 0
	for t := range ticker.C {
		fmt.Println("Tick at", t)
		if i >= len(artists) {
			break
		}

		artist := artists[i]
		i += 1

		mbid, err := api.artistMbid(artist)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		tags := api.topTagsMbid(mbid)

		var row []string
		row = append(row, artist)
		row = append(row, mbid)
		row = append(row, tags...)
		csv.Write(row)

		fmt.Printf("%q\t%s\n", artist, strings.Join(tags, "\t"))
		fmt.Printf("progress:\t%d of %d (%d%%)\n", i, artist_count, i/artist_count*100)
	}

}

func (l *lookupApi) ConnectWithConfig(configPath string) {
	configFile, err := os.Open(configPath)
	defer configFile.Close()
	if err != nil {
		panic(fmt.Errorf("Could not load config file %s", configPath))
	}

	var config lastfmConfig
	json.NewDecoder(configFile).Decode(&config)

	l.api = lastfm.New(config.Key, config.Secret)
}

func (l *lookupApi) artistMbid(artist string) (string, error) {
	result, err := l.api.Artist.Search(lastfm.P{"artist": artist})

	if err != nil {
		return "", err
	}
	if result.TotalResults <= 0 {
		return "", fmt.Errorf("Artist not found '%s'", artist)
	}

	mbid := result.ArtistMatches[0].Mbid
	return mbid, nil

}

func (l *lookupApi) topTagsMbid(mbid string) []string {
	var tags []string
	result, err := l.api.Artist.GetTopTags(lastfm.P{"mbid": mbid})
	if err != nil {
		fmt.Printf("%v\n", err)
		return tags
	}
	for _, tag := range result.Tags {
		tags = append(tags, tag.Name)
	}

	return tags
}
