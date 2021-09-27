package pmtools

import "net/http"

func (p *pmtools) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p.render(w, r, "home.page.tmpl", &templateData{})
	}
}
