package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.Handle("GET /snippet/create", app.requireAuthentication(http.HandlerFunc(app.snippetCreate)))
	mux.Handle("POST /snippet/create", app.requireAuthentication(http.HandlerFunc(app.snippetCreatePost)))
	mux.HandleFunc("GET /user/signup", app.userSignUp)
	mux.HandleFunc("POST /user/signup", app.userSignUpPost)
	mux.HandleFunc("GET /user/login", app.userLogIn)
	mux.HandleFunc("POST /user/login", app.userLogInPost)
	mux.Handle("POST /user/logout", app.requireAuthentication(http.HandlerFunc(app.userLogoutPost)))

	return app.logRequest(setCommonHeaders(app.sessionManager.LoadAndSave(app.authenticate(mux))))
}
