package middleware

import (
	"net/http"
	"strings"
)

// Heartbeat logs incoming requests on global HTTP handler
func Heartbeat(endpoint string) func(http.Handler) http.Handler {
	f := func(next http.Handler) http.Handler {
		fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if (r.Method == "GET" || r.Method == "HEAD") && strings.EqualFold(r.URL.Path, endpoint) {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("."))
				return
			}
			next.ServeHTTP(w, r)
		})
		return fn
	}
	return f
}
