package rest

// MiddlewareFilter defines a class of functions to filter middlewares.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router/rest instead.
type MiddlewareFilter func(middlewares ...Middleware) []Middleware

// NilMiddleware excludes nil middlewares from the list.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router/rest instead.
func NilMiddlewareFilter(middlewares ...Middleware) []Middleware {
	filtered := middlewares[:0]
	for _, middleware := range middlewares {
		if middleware != nil {
			filtered = append(filtered, middleware)
		}
	}
	return filtered
}
