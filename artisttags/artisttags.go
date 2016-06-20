package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

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

	api := lookupApi{}
	api.ConnectWithConfig("lastfm.json")

	artists := []string{"Radiohead", "Four tet", "Beatles"}

	ticker := time.NewTicker(time.Millisecond * 201)
	i := 0
	for t := range ticker.C {
		fmt.Println("Tick at", t)
		if i >= len(artists) {
			break
		}

		artist := artists[i]
		tags := api.topTags(artist)
		fmt.Printf("%q\t%s\n", artist, strings.Join(tags, "\t"))
		i += 1
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

func (l *lookupApi) artistMbid(artist string) string {
	result, err := l.api.Artist.Search(lastfm.P{"artist": artist})

	if err != nil {
		fmt.Println("%v", err)
		return ""
	}
	if result.TotalResults <= 0 {
		fmt.Println("Artist not found %s", artist)
		return ""
	}

	mbid := result.ArtistMatches[0].Mbid
	return mbid

}

func (l *lookupApi) topTags(artist string) []string {
	var tags []string
	result, err := l.api.Artist.GetTopTags(lastfm.P{"mbid": l.artistMbid(artist)})
	if err != nil {
		fmt.Println("%v", err)
	}
	for _, tag := range result.Tags {
		tags = append(tags, tag.Name)
	}
	return tags
}
