package pmtools

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (p *pmtools) routes() http.Handler {
	router := pat.New()

	router.Get("/", p.handleHome())

	fileserver := http.FileServer(http.Dir("./web/static/"))
	router.Get("/static/", http.StripPrefix("/static", fileserver))

	return router
}
