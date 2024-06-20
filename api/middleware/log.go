package middleware

import (
	"llm-manager/internal/log"
	"net/http"
)

// Logger prints an info  that includes remote addr, path and method
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.LoggerInstance.Logger.Printf("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
