package dj

import (
	"fmt"
	"log"
	"rfk/karma"
	"rfk/library"
	"rfk/observer"
	"sort"
)

func karmaSong() func() (library.Song, error) {
	log.Println("karmaSong")

	var impressionsByKey map[int][]string = invertMap(karma.SongKarma)
	impressionKeys := keys(impressionsByKey)
	sort.Ints(impressionKeys)

	log.Println("impressions prepped")
	songCh := make(chan *library.Song)
	observer.Observe("library.loaded", func(msg interface{}) {
		go func() {
			for i := len(impressionKeys) - 1; i >= 0; i-- {
				log.Printf("impression key: %v", impressionKeys[i])
				for _, hash := range impressionsByKey[impressionKeys[i]] {
					log.Printf("impression hash: %q", hash)
					song, err := library.ByHash(hash)
					if err != nil {
						continue
					}
					log.Printf("karmaSong karma:%d song:%v", impressionKeys[i], song)
					songCh <- song
				}
			}
			// after we runout, return nothing forever
			for {
				songCh <- nil
			}
		}()
	})
	return func() (library.Song, error) {
		nextSong := <-songCh
		if nextSong == nil {
			return library.Song{}, fmt.Errorf("NoSong")
		} else {
			return *nextSong, nil
		}
	}
}

func keys(hash map[int][]string) []int {
	var out []int
	for k, _ := range hash {
		out = append(out, k)
	}
	return out
}

func values(hash map[string]int) []int {
	var out []int
	for _, v := range hash {
		out = append(out, v)
	}
	return out
}

func uniq(arr []int) []int {
	var out []int
	set := make(map[int]bool)
	for _, v := range arr {
		if !set[v] {
			set[v] = true
			out = append(out, v)
		}
	}
	return out
}

func invertMap(hash map[string]int) map[int][]string {
	out := make(map[int][]string)
	for k, v := range hash {
		out[v] = append(out[v], k)
	}
	return out
}
