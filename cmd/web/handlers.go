package main

import (
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.html", nil)
}

func (app *application) codeView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Malformed ID", http.StatusBadRequest)
		return
	}

	app.render(w, r, "view.html", id)
}
