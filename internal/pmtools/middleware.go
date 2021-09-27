package pmtools

import (
	"fmt"
	"net/http"
)

func (p *pmtools) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				p.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (p *pmtools) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.logger.PrintInfo("request received", map[string]string{"method": r.Method, "url": r.URL.RequestURI()})
		next.ServeHTTP(w, r)
	})
}
