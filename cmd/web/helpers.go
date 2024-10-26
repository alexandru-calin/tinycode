package main

import (
	"html/template"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, files []string, data any) {
	tpl, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
