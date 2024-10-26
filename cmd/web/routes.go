package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /code/view/{id}", app.codeView)

	return app.logRequest(setCommonHeaders(mux))
}
