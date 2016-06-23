package main

import (
	"fmt"
	"sort"
)

// ------------------------------------------------

type Impression struct {
	SongHash   string
	Impression int
	State      string
}

type impressionList []Impression

func (s impressionList) Len() int           { return len(s) }
func (s impressionList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s impressionList) Less(i, j int) bool { return s[i].Impression < s[j].Impression }
func (s impressionList) IndexOf(target Impression) int {
	for i, impression := range s {
		if impression.SongHash == target.SongHash {
			return i
		}
	}
	return -1
}

func (self impressionList) AddUpdate(impression Impression) impressionList {
	if self == nil {
		impressions := make(impressionList, 0)
		self = impressions
	}
	var item Impression
	i := self.IndexOf(impression)
	if i >= 0 {
		fmt.Printf("TimeKm Update\n")
		item = self[i]
		item.Impression = item.Impression + impression.Impression
		self[i] = item
		return self
	} else {
		fmt.Printf("TimeKm Push\n")
		return append(self, impression)
	}
}

// ------------------------------------------------

type KarmaView interface {
	RelevantSongs() []Impression
	// internally update current state
}

type KarmaWriter interface {
	Impress(Impression)
	// internally knows about state
}

// ------------------------------------------------

type WeatherKm struct {
	impressions impressionList
}

func (self *WeatherKm) Impress(impression Impression) {
	fmt.Printf("WeatherKm Impress song: %s => %d\n", impression.SongHash, impression.Impression)
	self.impressions = self.impressions.AddUpdate(impression)
}

func (self *WeatherKm) RelevantSongs() []Impression {
	sort.Sort(self.impressions)
	return self.impressions
}

// ------------------------------------------------

type TimeKm struct {
	impressions impressionList
}

func (self *TimeKm) Impress(impression Impression) {
	impression.Impression *= 2 // XXX just to be different, debug
	fmt.Printf("TimeKm Impress song: %s => %d\n", impression.SongHash, impression.Impression)
	self.impressions = self.impressions.AddUpdate(impression)
}

func (self *TimeKm) RelevantSongs() []Impression {
	sort.Sort(self.impressions)
	return self.impressions
}

// ------------------------------------------------

type KarmaMultiWriter struct {
	writers []KarmaWriter
}

func (self *KarmaMultiWriter) RegisterKarmaWriter(writer KarmaWriter) {
	self.writers = append(self.writers, writer)
}
func (self *KarmaMultiWriter) Impress(impression Impression) {
	for _, writer := range self.writers {
		writer.Impress(impression)
	}
}

// ------------------------------------------------

type KarmaSpreader struct {
	destination KarmaWriter
}

func (k *KarmaSpreader) Impress(i Impression) {
	// pass along the original impression
	k.destination.Impress(i)

	// constructed with a SongGraph (or way to spread from one song to several)
	songHashes := k.getRelatedSongs(i.SongHash)
	for _, hash := range songHashes {
		// divide by two is fine for depth=1
		// for greater depth divide by more (e.g. impression/(depth+1))
		k.destination.Impress(Impression{SongHash: hash, Impression: i.Impression / 2})
	}
}
func (k *KarmaSpreader) getRelatedSongs(song string) []string {
	// return possibly related songids from the song.
	// use the lastfm artist tags
	var songs []string
	return songs
}

// ------------------------------------------------

type DJ struct {
	karmaViews []KarmaView
}

func (dj *DJ) RegisterKarmaView(k KarmaView) {
	dj.karmaViews = append(dj.karmaViews, k)
}

func (dj *DJ) relevantSongs() []Impression {
	var merged impressionList
	for _, kv := range dj.karmaViews {
		for _, impression := range kv.RelevantSongs() {
			merged = merged.AddUpdate(impression)
		}
	}
	return merged
}

// ------------------------------------------------

type ImpressionLog struct {
	impressionWriters       []KarmaWriter
	impressionSpreadWriters []KarmaWriter
}

func NewImpressionLog() *ImpressionLog {
	return &ImpressionLog{
		impressionWriters:       make([]KarmaWriter, 0),
		impressionSpreadWriters: make([]KarmaWriter, 0),
	}
}

func (self *ImpressionLog) Impress(impression Impression) {
	// direct
	for _, writer := range self.impressionWriters {
		writer.Impress(impression)
	}

	// spread
	for _, writer := range self.impressionSpreadWriters {
		fmt.Println(writer)
		writer.Impress(impression)
	}
	for _, spreadImpression := range self.getSpreadImpressions(impression) {
		for _, writer := range self.impressionSpreadWriters {
			writer.Impress(spreadImpression)
		}
	}

}

func (self *ImpressionLog) getSpreadImpressions(impression Impression) []Impression {
	var spreadImpressions []Impression
	return spreadImpressions
}

func (self *ImpressionLog) RegisterSpreadKarmaWriter(k KarmaWriter) {
	fmt.Println(k)
	self.impressionSpreadWriters = append(self.impressionSpreadWriters, k)
}

func main() {
	fmt.Println("Stub")

	timeKm := &TimeKm{}
	weatherKm := &WeatherKm{}

	impLog := NewImpressionLog()
	//impLog.RegisterKarmaWriter(fileWriter)
	impLog.RegisterSpreadKarmaWriter(weatherKm)
	impLog.RegisterSpreadKarmaWriter(timeKm)

	fmt.Println(impLog)

	dj := DJ{}
	dj.RegisterKarmaView(weatherKm)
	dj.RegisterKarmaView(timeKm)

	impressions := []Impression{
		Impression{SongHash: "a", Impression: 1},
		Impression{SongHash: "b", Impression: 1},
		Impression{SongHash: "c", Impression: -1},
		Impression{SongHash: "a", Impression: 1},
		Impression{SongHash: "b", Impression: -1},
		Impression{SongHash: "c", Impression: -1},
	}
	for _, impression := range impressions {
		fmt.Printf("outer Impress: %v\n", impression)
		impLog.Impress(impression)
	}

	for _, impression := range dj.relevantSongs() {
		fmt.Printf("merged: %v\n", impression)
	}

}
