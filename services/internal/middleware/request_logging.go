package middleware

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

// RequestLogging middleware to log incoming requests on global HTTP handler
func RequestLogging(logger *zap.Logger) Middleware {
	f := func(next http.Handler) http.Handler {
		fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Context().Value("")
			start := time.Now()

			next.ServeHTTP(w, r)
			// Log the request details
			logger.Info(
				"request",
				zap.String("method", r.Method),
				zap.String("ip", r.RemoteAddr),
				zap.String("url", r.RequestURI),
				zap.Duration("duration", time.Since(start)),
			)
		})
		return fn
	}
	return f
}
