package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/andynu/rfk/server/api"
	"github.com/andynu/rfk/server/library"
	"github.com/andynu/rfk/server/player"
)

func Listener() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/status", playerStatusHandler)
	http.HandleFunc("/next", nextHandler)
	http.HandleFunc("/skip", skipHandler)
	http.HandleFunc("/reward", rewardHandler)
	http.HandleFunc("/playpause", playPauseHandler)

	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/searchRequest", searchRequestHandler)
	http.HandleFunc("/searchUnrequest", searchUnrequestHandler)
	http.HandleFunc("/request", requestHandler)
	http.HandleFunc("/unrequest", unrequestHandler)
	http.HandleFunc("/requests", requestsHandler)
	http.HandleFunc("/clearRequests", clearRequestsHandler)

	http.HandleFunc("/stream", streamHandler)

	err := http.ListenAndServe(":7778", nil)
	if err != nil {
		panic(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexHTML)
}

func nextHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	api.SkipNoPunish()
	fmt.Fprintf(w, "ok")
}

func playPauseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	api.PlayPause()
	fmt.Fprintf(w, "ok")
}

func skipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	api.Skip()
	fmt.Fprintf(w, "ok")
}

func rewardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	api.Reward()
	fmt.Fprintf(w, "ok")
}

func playerStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/javascript")
	fmt.Fprintf(w, toJSON(api.PlayerStatus()))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/javascript")
	hash := r.URL.Query().Get("hash")
	var hashes []string
	hashes = append(hashes, hash)
	api.Request(hashes)
	fmt.Fprintf(w, toJSON(api.Requests()))
}

func unrequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/javascript")
	hash := r.URL.Query().Get("hash")
	var hashes []string
	hashes = append(hashes, hash)
	api.Unrequest(hashes)
	fmt.Fprintf(w, toJSON(api.Requests()))
}

func requestsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/javascript")
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, toJSON(api.Requests()))
	case "DELETE":
		api.ClearRequests()
		fmt.Fprintf(w, toJSON(api.Requests()))
	}
}

func clearRequestsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/javascript")
	api.ClearRequests()
	fmt.Fprintf(w, toJSON(api.Requests()))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/javascript")

	term := r.URL.Query().Get("term")
	offsetParam := r.URL.Query().Get("offset")
	limitParam := r.URL.Query().Get("limit")

	pagedSongs := pagedSongs(api.Search(term), limitParam, offsetParam)

	fmt.Fprintf(w, toJSON(pagedSongs))
}

func searchRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/javascript")

	term := r.URL.Query().Get("term")

	api.SearchRequest(term)

	fmt.Fprintf(w, toJSON(true))
}

func searchUnrequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/javascript")

	term := r.URL.Query().Get("term")

	api.SearchUnrequest(term)

	fmt.Fprintf(w, toJSON(true))
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

func pagedSongs(objects []*library.Song, limitStr string, offsetStr string) []*library.Song {
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 1000
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}
	startIdx := offset * limit
	endIdx := (offset + 1) * limit
	if endIdx > len(objects) {
		endIdx = len(objects)
	}
	return objects[startIdx:endIdx]
}
