package middleware

import (
	"net/http"
	"strings"
)

// Heartbeat logs incoming requests on global HTTP handler
func Heartbeat(next http.Handler, endpoint string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == "GET" || r.Method == "HEAD") && strings.EqualFold(r.URL.Path, endpoint) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("."))
			return
		}

		next.ServeHTTP(w, r)
	})
}
