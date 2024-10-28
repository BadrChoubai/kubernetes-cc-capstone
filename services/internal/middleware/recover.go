package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/badrchoubai/services/internal/encoding"
	"github.com/badrchoubai/services/internal/observability/logging"
)

type errorResponse struct {
	ApplicationError map[string]any `json:"applicationError"`
}

func Recover(logger *logging.Logger) func(next http.Handler) http.Handler {
	encoderDecoder := encoding.NewEncoderDecoder()

	f := func(next http.Handler) http.Handler {
		fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := &errorResponse{}

			defer func() {
				if err := recover(); err != nil {
					w.Header().Set("Connection", "close")

					var errMsg string
					switch e := err.(type) {
					case error:
						errMsg = e.Error()
					case string:
						errMsg = e
					default:
						errMsg = "Unknown error occurred"
					}

					errParts := strings.SplitN(errMsg, ":", 2)
					errMap := map[string]any{}
					if len(errParts) == 2 {
						errMap[errParts[0]] = strings.TrimSpace(errParts[1])
					} else {
						errMap["error"] = errMsg // Fallback if ":" isn't in the error message
					}

					response.ApplicationError = errMap

					logger.Error("application error", errors.New(errMsg))
					_ = encoderDecoder.EncodeResponse(w, http.StatusInternalServerError, response)
				}
			}()
			next.ServeHTTP(w, r)
		})
		return fn
	}
	return f
}
