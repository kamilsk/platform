package http

import "net/http"

// Router is an HTTP request multiplexer.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router.Interface instead.
type Router interface {
	// Handle registers the handler for the given pattern.
	Handle(string, http.Handler)
	// HandleFunc registers the handler function for the given pattern.
	HandleFunc(string, func(http.ResponseWriter, *http.Request))
}

// Endpoint represents HTTP listener that can register its routes.
//
// Deprecated: doesn't have canonical implementation, use
// go.octolab.org/toolkit/protocol/http/router instead.
type Endpoint interface {
	// Register registers itself in the router.
	Register(Router)
}
