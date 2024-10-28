package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, page string, data any) {
	tpl, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	err := tpl.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	buf.WriteTo(w)
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}
