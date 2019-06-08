package rest

import "net/http"

// Handler is a http handler with a specified path.
type Handler func() (path string, handler http.Handler)

// HandlerFunc is a http handler with a specified path.
type HandlerFunc func() (path string, handler http.HandlerFunc)

// PackedHandler is a http handler with a specified http method and path.
type PackedHandler func() (method string, path string, handler http.Handler)

// PackedHandlerFunc is a http handler with a specified http method and path.
type PackedHandlerFunc func() (method string, path string, handler http.HandlerFunc)
