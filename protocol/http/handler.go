package http

import "net/http"

// Handler is a http handler with a specified path.
type Handler func() (path string, handler http.Handler)

// HandlerFunc is a http handler with a specified path.
type HandlerFunc func() (path string, handler http.HandlerFunc)
