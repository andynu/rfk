package rest

import (
	"fmt"
	"net/http"
)

func RESTListener() {
	go func() {
		http.HandleFunc("/", root_handler)
		err := http.ListenAndServe(":7778", nil)
		if err != nil {
			panic(err)
		}
	}()
}

func root_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RFK!")
}
