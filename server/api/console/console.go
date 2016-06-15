package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andynu/rfk/server/api"
	"github.com/andynu/rfk/server/library"
)

func Listener() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		args := strings.Split(text, " ")
		command := args[0]
		switch command {
		case "n":
			log.Printf("cmd: n - Next (no punishment)")
			api.SkipNoPunish()

		case "s":
			log.Printf("cmd: s - Skip")
			api.Skip()

		case "p":
			log.Printf("cmd: p - PlayPause")
			api.PlayPause()

		case "r":
			log.Printf("cmd: r - Reward")
			api.Reward()

		case "search":
			term := strings.Join(args[1:], " ")
			log.Printf("cmd: search - %q", term)
			songs := api.Search(term)
			printSearchResults(term, songs)

		case "request":
			term := strings.Join(args[1:], " ")
			log.Printf("cmd: request - %q", term)
			songs := api.SearchRequest(term)
			fmt.Printf("Requested %d songs\n", len(songs))

		case "taglast":
			tag := strings.Join(args[1:], " ")
			log.Printf("cmd: tag last song - %q", tag)
			api.TagLastSong(tag)

		case "tag":
			tag := strings.Join(args[1:], " ")
			log.Printf("cmd: tag current song - %q", tag)
			api.TagCurrentSong(tag)

		case "clear":
			log.Printf("cmd: clear request")
			api.ClearRequests()
			fmt.Printf("Requests Cleared\n")

		default:
			log.Printf("cmd: %q - unknown command", text)
		}
	}
}

func printSearchResults(term string, songs []*library.Song) {
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
