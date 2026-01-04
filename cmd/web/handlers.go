package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) createView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create"))
}

func (app *application) createPOST(w http.ResponseWriter, r *http.Request) {
	topic := "leetcode"
	duration := 60
	notes := "Studied some traversal stuff on trees"

	id, err := app.sessions.InsertSession(topic, duration, notes)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/study/view?id=%d", id), http.StatusSeeOther)
}

func (app *application) studyView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Getting study with id %d...", id)
}

func (app *application) viewTopicStats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("View stats"))
}
