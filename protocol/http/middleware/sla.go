package middleware

import (
	"context"
	"net/http"
	"time"

	platform "github.com/kamilsk/platform/protocol/http"
)

// SLA returns a middleware to limit the lifetime of a request context.
func SLA(fallback time.Duration, correction float64) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			ctx, cancel := context.WithTimeout(req.Context(), platform.Header(req.Header).Timeout(fallback, correction))
			defer cancel()

			handler.ServeHTTP(resp, req.WithContext(ctx))
		})
	}
}
