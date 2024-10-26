package observability

import (
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/badrchoubai/services/internal/observability/logging"
)

// RequestLoggingMiddleware logs incoming requests on global HTTP handler
func RequestLoggingMiddleware(logger *logging.Logger) func(http.Handler) http.Handler {
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
