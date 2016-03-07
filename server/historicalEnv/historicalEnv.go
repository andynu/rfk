package historicalEnv

import (
	"log"
	"strconv"
	"time"

	"github.com/andynu/rfk/server/library"
)

type ValueImpression map[string]int
type TagValue map[string]ValueImpression
type SongEnv map[*library.Song]TagValue

// env = tag: value
var env map[string]string

// song: { tag: { value: impression } }
var songEnv SongEnv

func init() {
	env = make(map[string]string)
	songEnv = make(SongEnv)
}

func impress(song *library.Song, tag string, val string, impression int) {
	//log.Printf("%s\t%s\t%s\t%d\n", song.Hash, tag, val, impression)
	if _, ok := songEnv[song]; !ok {
		songEnv[song] = make(TagValue)
	}
	if _, ok := songEnv[song][tag]; !ok {
		songEnv[song][tag] = make(ValueImpression)
	}
	songEnv[song][tag][val] += impression
}

func updateTime(timestamp time.Time) {
	Update(timestamp, "hour", strconv.Itoa(timestamp.Hour()))
	Update(timestamp, "weekday", strconv.Itoa(int(timestamp.Weekday())))
}

func Update(timestamp time.Time, tag string, val string) {
	env[tag] = val
}

func Impress(timestamp time.Time, song *library.Song, impression int) {
	updateTime(timestamp)
	for k := range env {
		impress(song, env[k], k, impression)
	}
}

func Print() {
	for s := range songEnv {
		for t := range songEnv[s] {
			for v := range songEnv[s][t] {
				i := songEnv[s][t][v]
				log.Printf("%s\t%s\t%s\t%d\n", s, t, v, i)
			}
		}
	}
}
