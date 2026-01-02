package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Home"))
}

func createView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create"))
}

func createPOST(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create POST"))
}

func studyView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Getting study with id %d...", id)
}
