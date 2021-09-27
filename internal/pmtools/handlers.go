package pmtools

import (
	"fmt"
	"net/http"

	"github.com/pscompsci/pmtools/pkg/forms"
)

func (p *pmtools) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p.render(w, r, "home.page.tmpl", &templateData{})
	}
}

func (p *pmtools) handleNewSnapshot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p.render(w, r, "snapshot.form.page.tmpl", &templateData{MultipartForm: forms.NewMultipartForm()})
	}
}

func (p *pmtools) handleNewShapshotSubmit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			p.logger.PrintInfo(fmt.Sprintf("%s - %s - server error", r.URL, r.Method), nil)
			p.serverError(w, err)
			return
		}

		form := forms.NewMultipartForm()
		form.Required(r, "revenue", "credits", "gar")
		form.IsType(r, "revenue", ".csv")
		form.IsType(r, "credits", ".csv")
		form.IsType(r, "gar", ".csv")

		if !form.Valid() {
			p.render(w, r, "snapshot.form.page.tmpl", &templateData{MultipartForm: form})
			return
		}

		w.Write([]byte("Form submitted ok")) // continue from here
	}
}
