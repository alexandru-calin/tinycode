package main

import (
	"net/http"
)

func (app *application) render(w http.ResponseWriter, page string, data any) {
	tpl := app.templateCache[page]

	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
