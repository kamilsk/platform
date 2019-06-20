package rest

import "net/http"

// Handler is a http handler with a specified http method and path.
type Handler func() (method, path string, handler http.Handler)

// HandlerFunc is a http handler with a specified http method and path.
type HandlerFunc func() (method, path string, handler http.HandlerFunc)
