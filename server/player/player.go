// Plays mp3 files
package player

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/andynu/rfk/server/dj"
	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/observer"

	"github.com/dhowden/tag"
)

const playerBin string = "mpg123"

var CurrentSong *library.Song
var LastSong *library.Song

var playerCmd *exec.Cmd

var playing_mu sync.Mutex
var playing bool = true

var sleepPlayer bool = false

func PlaySongs() {
	//var wait_for_dj_songs chan bool
	//observer.Observe("karma.loaded", func(msg interface{}) {
	//	wait_for_dj_songs <- true
	//})
	//<-wait_for_dj_songs

	err_count := 0
	for {
		if !IsPlaying() {
			time.Sleep(1 * time.Second)
			continue
		}

		song, serr := dj.NextSong()
		if serr != nil {
			log.Printf("rfk: %v", serr)
			err_count += 1
			if err_count > 5 {
				panic("DJFail FATAL!")
			}
			continue
		}
		err_count = 0

		perr := Play(song)
		if perr != nil {
			log.Printf("rfk: %v", perr)
			time.Sleep(1 * time.Second)
		}
	}
}

func Silence() {
	sleepPlayer = true
}

func PlayPauseState() string {
	playing_mu.Lock()
	defer playing_mu.Unlock()
	if playing {
		return "playing"
	}
	return "paused"
}

func Play(song library.Song) error {
	LastSong = CurrentSong
	CurrentSong = &song
	log.Printf("player: playing %q (%f)", song.Path, song.Rank)
	logMetadata(song.Path)

	if sleepPlayer {
		playerCmd = exec.Command("sleep", "5")
	} else {
		playerCmd = exec.Command(playerBin, song.Path)
	}
	err := playerCmd.Run()
	if err != nil {
		return fmt.Errorf("player: %v", err)
	}
	observer.Notify("player.played", &song)
	return nil
}

func Stop() error {
	if playerCmd != nil {
		return playerCmd.Process.Kill()
	}
	return nil
}

func Skip() error {
	observer.Notify("player.skip", CurrentSong)
	return Stop()
}

func IsPlaying() bool {
	playing_mu.Lock()
	isPlaying := playing
	playing_mu.Unlock()
	return isPlaying
}

func TogglePlayPause() {
	playing_mu.Lock()
	playing = !playing
	playing_mu.Unlock()
	if !playing {
		Stop()
	}
}

func Pause() {
	playing_mu.Lock()
	playing = false
	playing_mu.Unlock()
	Stop()
}

func Unpause() {
	playing_mu.Lock()
	playing = true
	playing_mu.Unlock()
}

func logMetadata(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("Could not load metadata: %q", err)
		return
	}
	defer f.Close()

	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Printf("Could not load metadata: %q", err)
		return
	}

	log.Printf("\tTitle: %q", m.Title())
	log.Printf("\tAlbum: %q", m.Album())
	log.Printf("\tArtist: %q", m.Artist())
}
