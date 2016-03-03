package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/config"
)

func ensureConfig(configPath string){
	configPath = config.DefaultConfigPath(configPath)
	config.CreateDefaultConfig(configPath)
}

func ensureDataPath(dataPath string) {
	dataPath = config.DefaultDataPath(dataPath)
	config.CreateDataPath(dataPath)
}


func ensureSongs(){
	songsPath := path.Join(config.DataPath, "songs.txt")
	songHashesPath := path.Join(config.DataPath, "song_hashes.txt")
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



func gets() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return text
}
