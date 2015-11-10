package main

import (
	"log"
	"os/exec"
)

func main() {
	log.Println("Rfk v3 startup")

	playerBin := "mpg123"
	mp3Path := "02-Bodysnatchers.mp3"
  out, err := exec.Command(playerBin, mp3Path).Output()
  if err != nil {
    log.Fatal(err)
  }
  log.Printf(string(out))
}

