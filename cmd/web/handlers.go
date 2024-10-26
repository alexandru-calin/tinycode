package main

import (
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{"./ui/html/base.html", "./ui/html/pages/home.html"}
	app.render(w, r, files, nil)
}

func (app *application) codeView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Malformed ID", http.StatusBadRequest)
		return
	}

	files := []string{"./ui/html/base.html", "./ui/html/pages/view.html"}
	app.render(w, r, files, id)
}
