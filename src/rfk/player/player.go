// Plays mp3 files
package player

import (
	"log"
	"os/exec"
)

type songPath string

const playerBin string = "mpg123"

func Play(path string) {
  log.Printf("player: playing %q", path)
	out, err := exec.Command(playerBin, path).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(string(out))
}
