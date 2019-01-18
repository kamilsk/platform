package http

import "net/http"

// Router is an HTTP request multiplexer.
type Router interface {
	// Handle registers the handler for the given pattern.
	Handle(string, http.Handler)
	// HandleFunc registers the handler function for the given pattern.
	HandleFunc(string, func(http.ResponseWriter, *http.Request))
}

// Endpoint represents HTTP listener that can register its routes.
type Endpoint interface {
	// Register registers itself in the router.
	Register(Router)
}
