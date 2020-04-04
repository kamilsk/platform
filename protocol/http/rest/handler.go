package rest

import "net/http"

// Handler is a http handler with a specified http method and path.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router/rest.Handler instead.
type Handler func() (method, path string, handler http.Handler)

// HandlerFunc is a http handler with a specified http method and path.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router/rest.HandlerFunc instead.
type HandlerFunc func() (method, path string, handler http.HandlerFunc)
