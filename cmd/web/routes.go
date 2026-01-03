package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("POST /study/create", app.createPOST)
	mux.HandleFunc("GET /study/create", app.createView)
	mux.HandleFunc("/study/view", app.studyView)

	return mux
}
