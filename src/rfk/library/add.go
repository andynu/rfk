package library

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"sort"
)

// Adds mp3s to the library (songs.txt) from the specified paths.
//
// A 100% go version of `find <path> -name '*.mp3' > songs.txt`
func AddPaths(paths []string) {
	f, err := os.OpenFile(songsPath, os.O_APPEND|os.O_WRONLY, 0600)
	panicOnErr(err)
	mp3Count := 0

	for _, path := range paths {
		log.Printf("adding path=%q", path)
		walkByExt(path, ".mp3", func(mp3path string) {
			mp3Count++
			f.WriteString(mp3path + "\n")
		})
	}
	f.Close()

	uniqueSortFileLines(songsPath)

	log.Printf("found %d songs", mp3Count)
}

func uniqueSortFileLines(filePath string) {
	uniqueLines := make(map[string]bool)
	var lines []string

	// read
	inFile, err := os.Open(filePath)
	panicOnErr(err)
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		if !uniqueLines[line] {
			lines = append(lines, line)
			uniqueLines[scanner.Text()] = true
		}
	}
	inFile.Close()

	sort.Strings(lines)

	// write
	outFile, err := os.Create(filePath)
	panicOnErr(err)
	for _, line := range lines {
		outFile.WriteString(line + "\n")
	}
	outFile.Close()
}

func walkByExt(rootPath string, ext string, task func(string)) {

	visit := func(path string, file os.FileInfo, err error) error {
		if err != nil {
			log.Println("walk err: %v", err)
		}
		//log.Printf("debug: %q %q", path, file.Name())
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ext {
				task(path)
			}
		}
		return nil
	}
	err := filepath.Walk(rootPath, visit)
	if err != nil {
		log.Printf("err: %v", err)
	}
}
