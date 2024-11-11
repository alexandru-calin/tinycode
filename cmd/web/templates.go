package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/alexandru-calin/tinycode/internal/models"
	"github.com/alexandru-calin/tinycode/ui"
)

type templateData struct {
	Snippet         models.Snippet
	Snippets        []models.Snippet
	Form            any
	Toast           string
	IsAuthenticated bool
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		patterns := []string{"html/base.html", page}

		tpl, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
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
