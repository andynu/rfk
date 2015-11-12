// Plays mp3 files
package player

import (
	"log"
	"os"
	"os/exec"

	"github.com/dhowden/tag"
)

const playerBin string = "mpg123"

func Play(path string) {
	log.Printf("player: playing %q", path)
	logMetadata(path)

	out, err := exec.Command(playerBin, path).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(string(out))
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
