package pmtools

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func (p *pmtools) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	log.Println(trace)
}

func (p *pmtools) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (p *pmtools) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	return td
}

func (p *pmtools) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	templates, ok := p.templates[name]
	if !ok {
		p.serverError(w, fmt.Errorf("the template '%s' doe not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := templates.Execute(buf, p.addDefaultData(td, r))
	if err != nil {
		p.serverError(w, err)
	}
	buf.WriteTo(w)
}
