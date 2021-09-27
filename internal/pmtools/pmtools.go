package pmtools

import (
	"html/template"
	"os"
	"sync"

	"github.com/pscompsci/pmtools/internal/cache"
	"github.com/pscompsci/pmtools/pkg/log"
)

type pmtools struct {
	config    *Config
	logger    *log.Logger
	cache     *cache.Cache
	templates map[string]*template.Template
	wg        sync.WaitGroup
}

func New(cfg *Config) *pmtools {
	logger := log.New(os.Stdout, log.LevelInfo)
	templates, err := newTemplateCache("./web/templates/")
	if err != nil {
		logger.PrintFatal("failed to load the template cache", map[string]string{"error": err.Error()})
	}

	return &pmtools{
		config:    cfg,
		logger:    logger,
		cache:     cache.New(),
		templates: templates,
	}
}

func (p *pmtools) Run() error {
	return p.serve()
}
