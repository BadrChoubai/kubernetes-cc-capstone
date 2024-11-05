// Package middleware provides various middleware functions that can be applied to HTTP servers.
// Middleware is a function that wraps a http.Handler to perform operations such as logging,
// authentication, request/response modification, and error handling.
package middleware

import "net/http"

// Middleware type
type Middleware func(http.Handler) http.Handler
