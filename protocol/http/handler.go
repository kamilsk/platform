package http

import "net/http"

// Handler is a http handler with a specified path.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router.Handler instead.
type Handler func() (path string, handler http.Handler)

// HandlerFunc is a http handler with a specified path.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router.HandlerFunc instead.
type HandlerFunc func() (path string, handler http.HandlerFunc)
