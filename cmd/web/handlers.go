package main

import (
	"errors"
	"fmt"

	//"html/template"
	"net/http"
	"strconv"
	"study-tracker-nickh/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	sessions, err := app.sessions.GetLatestSessions()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, session := range sessions {
		fmt.Fprintf(w, "%+v\n", session)
	}

	// files := []string{
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/pages/home.tmpl",
	// 	"./ui/html/partials/nav.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)

	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
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

	session, err := app.sessions.GetSession(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}

		return
	}
	fmt.Fprintf(w, "%+v", session)
}

func (app *application) viewTopicStats(w http.ResponseWriter, r *http.Request) {
	stats, err := app.stats.GetSessionStats()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, stat := range stats {
		fmt.Fprintf(w, "%+v\n", stat)
	}
}
