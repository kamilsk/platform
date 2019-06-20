package http

import "net/http"

// Handler is a http handler with a specified path.
type Handler func() (path string, handler http.Handler)

// HandlerFunc is a http handler with a specified path.
type HandlerFunc func() (path string, handler http.HandlerFunc)

// PackHandler wraps the path and handler.
func PackHandler(path string, handler http.Handler) Handler {
	return func() (string, http.Handler) { return path, handler }
}

// PackHandlerFunc wraps the path and handler.
func PackHandlerFunc(path string, handler http.HandlerFunc) HandlerFunc {
	return func() (string, http.HandlerFunc) { return path, handler }
}
