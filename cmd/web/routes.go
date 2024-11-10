package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	mux.HandleFunc("GET /user/signup", app.userSignUp)
	mux.HandleFunc("POST /user/signup", app.userSignUpPost)
	mux.HandleFunc("GET /user/login", app.userLogIn)
	mux.HandleFunc("POST /user/login", app.userLogInPost)
	mux.HandleFunc("POST /user/logout", app.userLogoutPost)

	return app.logRequest(setCommonHeaders(app.sessionManager.LoadAndSave(mux)))
}
