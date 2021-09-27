package pmtools

import (
	"html/template"
	"path/filepath"

	"github.com/pscompsci/pmtools/pkg/forms"
)

type templateData struct {
	MultipartForm *forms.MultipartForm
}

var functions = template.FuncMap{}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templates, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		templates, err = templates.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		templates, err = templates.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = templates
	}

	return cache, nil
}
