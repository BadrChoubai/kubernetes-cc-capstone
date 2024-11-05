package middleware

import (
	"net/http"
)

// Cors middleware
func Cors(enabled bool, trustedOrigins []string) Middleware {
	f := func(next http.Handler) http.Handler {
		fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if enabled {
				w.Header().Add("Vary", "Origin")
				origin := r.Header.Get("Origin")

				if origin != "" {
					for i := range trustedOrigins {
						if origin == trustedOrigins[i] {
							w.Header().Set("Access-Control-Allow-Origin", origin)

							if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
								w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, PUT, PATCH, DELETE")
								w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

								w.WriteHeader(http.StatusOK)
								return
							}

							break
						}
					}
				}
			}
			next.ServeHTTP(w, r)
		})
		return fn
	}
	return f
}
