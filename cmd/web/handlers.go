package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./ui/html/pages/home.html")
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (app *application) codeView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Malformed ID", http.StatusBadRequest)
		return
	}

	tpl, err := template.ParseFiles("./ui/html/pages/view.html")
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, id)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
