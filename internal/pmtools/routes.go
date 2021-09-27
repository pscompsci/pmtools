package pmtools

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (p *pmtools) routes() http.Handler {
	standardMiddleware := alice.New(p.recoverPanic, p.logRequest)

	router := pat.New()

	router.Get("/", p.handleHome())

	router.Get("/snapshot", p.handleNewSnapshot())
	router.Post("/snapshot", p.handleNewShapshotSubmit())

	fileserver := http.FileServer(http.Dir("./web/static/"))
	router.Get("/static/", http.StripPrefix("/static", fileserver))

	return standardMiddleware.Then(router)
}
