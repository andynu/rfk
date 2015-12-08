package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andynu/rfk/server/dj"
	"github.com/andynu/rfk/server/karma"
	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/player"
)

func InputListener() {
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, _ := reader.ReadString('\n')
			text = strings.TrimSuffix(text, "\n")
			args := strings.Split(text, " ")
			command := args[0]
			switch command {
			case "n":
				log.Printf("cmd: n - Stop")
				player.Stop()
			case "s":
				log.Printf("cmd: s - Skip")
				player.Skip()
			case "p":
				log.Printf("cmd: p - PlayPause")
				player.TogglePlayPause()
			case "r":
				log.Printf("cmd: r - Reward")
				karma.Log(player.CurrentSong, 1)
			case "search":
				term := strings.Join(args[1:], " ")
				log.Printf("cmd: search - %q", term)
				search(term)
			case "request":
				term := strings.Join(args[1:], " ")
				log.Printf("cmd: request - %q", term)
				request(term)
			case "taglast":
				tag := strings.Join(args[1:], " ")
				log.Printf("cmd: tag last song - %q", tag)
				tagLast(tag)
			case "tag":
				tag := strings.Join(args[1:], " ")
				log.Printf("cmd: tag current song - %q", tag)
				tagCurrent(tag)
			case "clear":
				log.Printf("cmd: clear request")
				clearRequests()
			default:
				log.Printf("cmd: %q - unknown command", text)
			}
		}
	}()
}

func search(term string) {
	songs := library.Search(term)
	fmt.Printf("\n")

	fmt.Printf("found %d songs for term %q\n", len(songs), term)
	limit := 10
	if len(songs) < limit {
		limit = len(songs)
	}

	for i, song := range songs[0:limit] {
		fmt.Printf("\t%d. %q\n", i, song.Path)
	}
}

func request(term string) {
	songs := library.Search(term)
	dj.Request(songs)
	fmt.Printf("Requested %d songs\n", len(songs))
}

func clearRequests() {
	dj.ClearRequests()
	fmt.Printf("Requests Cleared\n")
}

func tagCurrent(tag string) {
	karma.LogTag(player.CurrentSong, tag)
}

func tagLast(tag string) {
	karma.LogTag(player.LastSong, tag)
}
