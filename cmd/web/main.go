package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("POST /study/create", createPOST)
	mux.HandleFunc("GET /study/create", createView)
	mux.HandleFunc("/study/view", studyView)
	log.Println("Listening")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
