// Plays mp3 files
package player

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/observer"

	"github.com/dhowden/tag"
)

const playerBin string = "mpg123"

var CurrentSong library.Song

var playerCmd *exec.Cmd

var playing_mu sync.Mutex
var playing bool = true

func Play(song library.Song) error {
	CurrentSong = song
	log.Printf("player: playing %q", song.Path)
	logMetadata(song.Path)

	playerCmd = exec.Command(playerBin, song.Path)
	//playerCmd = exec.Command("sleep", "5")
	err := playerCmd.Run()
	if err != nil {
		return fmt.Errorf("player: %v", err)
	}
	observer.Notify("player.played", song)
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
