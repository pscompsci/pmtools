package pmtools

import (
	"html/template"
	"log"
	"sync"
)

type pmtools struct {
	config    *Config
	templates map[string]*template.Template
	wg        sync.WaitGroup
}

func New(cfg *Config) *pmtools {
	templates, err := newTemplateCache("./web/templates/")
	if err != nil {
		log.Fatal("could not parse templates")
	}

	return &pmtools{
		config:    cfg,
		templates: templates,
	}
}

func (p *pmtools) Run() error {
	return p.serve()
}
