// Plays mp3 files
package player

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"rfk/library"
	"rfk/observer"

	"github.com/dhowden/tag"
)

const playerBin string = "mpg123"

var CurrentSong library.Song
var playerCmd *exec.Cmd

func Play(song library.Song) error {
	CurrentSong = song
	log.Printf("player: playing %q", song.Path)
	logMetadata(song.Path)

	//playerCmd = exec.Command(playerBin, song.Path)
	playerCmd = exec.Command("sleep", "5")
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
