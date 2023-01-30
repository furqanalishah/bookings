package render

import (
	"bytes"
	"github.com/furqanalishah/bookings/pkg/config"
	"github.com/furqanalishah/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders an html template
func RenderTemplate(w http.ResponseWriter, t string, td *models.TemplateData) {
	var cache config.TemplateCache
	if app.UseCache {
		cache = app.TemplateCache
	} else {
		cache, _ = CreateTemplateCache()
	}

	tmpl, ok := cache[t]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	if err := tmpl.Execute(buf, td); err != nil {
		log.Fatal(err)
	}
	if _, err := buf.WriteTo(w); err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCache() (config.TemplateCache, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}

	return cache, nil
}
