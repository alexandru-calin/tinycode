package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/alexandru-calin/tinycode/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.html", nil)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Malformed ID", http.StatusBadRequest)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	app.render(w, r, "view.html", snippet.ID)
}
