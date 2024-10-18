package zap

import (
	"github.com/badrchoubai/services/internal/observability/logging"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// RequestLoggingMiddleware logs incoming requests on global HTTP handler
func RequestLoggingMiddleware(handler http.Handler, logger logging.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Context().Value("")
		start := time.Now()

		handler.ServeHTTP(w, r)
		// Log the request details
		logger.Info(
			"request",
			zap.String("method", r.Method),
			zap.String("ip", r.RemoteAddr),
			zap.String("url", r.RequestURI),
			zap.Duration("duration", time.Since(start)),
		)
	})
}
