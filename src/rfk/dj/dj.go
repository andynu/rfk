// Selects songs for a Player
package dj

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NextSong(songsTxt string) string {
	return randomSong(songsTxt)
}

func randomSong(songsTxt string) string {
	rand.Seed(time.Now().UnixNano())
	f, err := os.Open(songsTxt)
	check(err)
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines[rand.Intn(len(lines))]
}
