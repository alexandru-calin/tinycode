package main

import (
	"html/template"
	"path/filepath"
)

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{"./ui/html/base.html", page}

		tpl, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = tpl
	}

	return cache, nil
}
