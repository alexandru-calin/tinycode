package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/alexandru-calin/tinycode/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
	Form     any
	Toast    string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{"./ui/html/base.html", page}

		tpl, err := template.New(name).Funcs(functions).ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = tpl
	}

	return cache, nil
}

var functions = template.FuncMap{
	"humanReadableDate": humanReadableDate,
}

func humanReadableDate(t time.Time) string {
	return t.Format("02 January 2006 at 15:04")
}
