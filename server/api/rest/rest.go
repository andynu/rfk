package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/andynu/rfk/server/api"
	"github.com/andynu/rfk/server/player"
)

func RESTListener() {
	go func() {

		http.HandleFunc("/", rootHandler)
		http.HandleFunc("/playing", currentSongHandler)
		http.HandleFunc("/last_song", lastSongHandler)
		http.HandleFunc("/next", nextHandler)
		http.HandleFunc("/skip", skipHandler)
		http.HandleFunc("/playpause", playPauseHandler)

		http.HandleFunc("/stream", streamHandler)

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
	w.Header().Set("Content-type", "application/javascript")
	fmt.Fprintf(w, toJSON(player.CurrentSong))
}

func lastSongHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/javascript")
	fmt.Fprintf(w, toJSON(player.LastSong))
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	for {

		streamingSong := &player.CurrentSong

		streamBytes, err := ioutil.ReadFile(player.CurrentSong.Path)

		if err != nil {
			fmt.Println(err)
			return
		}

		b := bytes.NewBuffer(streamBytes)

		// stream straight to client(browser)
		w.Header().Set("Content-type", "audio/mpeg")

		if _, err := b.WriteTo(w); err != nil { // <----- here!
			fmt.Fprintf(w, "%s", err)
		}

		// wait for next song
		for {
			if &player.CurrentSong == streamingSong {
				time.Sleep(1 * time.Second)
			} else {
				break
			}
		}

		//w.Write([]byte("PDF Generated"))
	}
}

func toJSON(obj interface{}) string {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err) // Safe to return this err?
		return ""
	}
	return string(data)
}
