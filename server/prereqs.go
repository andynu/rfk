package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/andynu/rfk/server/library"
)

func checkPrereqs() {
	ensureBinaryExists("mpg123")

	scriptDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	configPath := path.Join(scriptDir, "config.json")
	if !pathExists(configPath) {
		fmt.Println(configPath)
		configExamplePath := path.Join(scriptDir, "config.json.example")
		cp(configExamplePath, configPath)
		log.Printf("Missing config: created default at %q", configPath)
	}

	dataPath := path.Join(scriptDir, "data")
	if !pathExists(dataPath) {
		log.Printf("Missing data dir: creating at %q", dataPath)
		os.Mkdir(dataPath, 0760)
	}

	songsPath := path.Join(scriptDir, "data", "songs.txt")
	songHashesPath := path.Join(scriptDir, "data", "song_hashes.txt")
	if !pathExists(songsPath) && !pathExists(songHashesPath) {
		log.Printf("No song indexes detected!")
		fmt.Printf("\nPlease provide folder with mp3s: ")
		musicPath := gets()
		for !pathExists(musicPath) {
			fmt.Printf("\nInvalid path. Please try again: ")
			musicPath = gets()
		}
		library.AddPaths([]string{musicPath})
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return (err == nil)
}

func ensureBinaryExists(executable string) {
	_, err := exec.LookPath(executable)
	if err != nil {
		panic(fmt.Errorf("prereq: %q [failed]: %q", executable, err))
	}
}

func cp(src, dst string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	d, err := os.Create(dst)
	if err != nil {
		return err
	}

	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}

func gets() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return text
}
