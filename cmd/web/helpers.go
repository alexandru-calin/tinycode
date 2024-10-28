package main

import (
	"net/http"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, page string, data any) {
	tpl := app.templateCache[page]

	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}
