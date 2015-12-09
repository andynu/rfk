package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andynu/rfk/server/api"
	"github.com/andynu/rfk/server/player"
)

func RESTListener() {
	go func() {

		http.HandleFunc("/", rootHandler)
		http.HandleFunc("/playing", currentSongHandler)
		http.HandleFunc("/next", nextHandler)
		http.HandleFunc("/skip", skipHandler)
		http.HandleFunc("/play", playPauseHandler)
		http.HandleFunc("/pause", playPauseHandler)

		err := http.ListenAndServe(":7778", nil)
		if err != nil {
			panic(err)
		}
	}()
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexHTML)
}

func nextHandler(w http.ResponseWriter, r *http.Request) {
	api.SkipNoPunish()
	fmt.Fprintf(w, "ok")
}

func playPauseHandler(w http.ResponseWriter, r *http.Request) {
	api.PlayPause()
	fmt.Fprintf(w, "ok")
}

func skipHandler(w http.ResponseWriter, r *http.Request) {
	api.Skip()
	fmt.Fprintf(w, "ok")
}

func rewardHandler(w http.ResponseWriter, r *http.Request) {
	api.Reward()
	fmt.Fprintf(w, "ok")
}

func currentSongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, toJSON(player.CurrentSong))
}

func toJSON(obj interface{}) string {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err) // Safe to return this err?
		return ""
	}
	return string(data)
}
