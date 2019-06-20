// Package rest provides handle types to describe a RESTful API.
//
//  mux := http.NewServeMux()
//  for _, handler := range routing.Handlers() {
//  	mux.Handle(handler())
//  }
//  http.ListenAndServe("localhost:8080", mux)
//
//  -- routing/routes.go --
//
//  func Handlers() []rest.Handler {
//  	return []rest.Handler{
//  		func() (string, http.Handler) { return "/ping", http.HandlerFunc(api.Pong) }
//  	}
//  }
//
//  -- api/handlers.go --
//
//  func Pong(rw http.ResponseWriter, req *http.Request) { _, _ = rw.Write([]byte("pong")) }
//
// It also friendly to advanced routers like github.com/go-chi/chi.
//
//  mux := http.NewServeMux()
//  mux.Handle(
//  	routing.V1("/v1/",
//  		chi.PackHandler(http.MethodGet, v1.Pong("/ping")),
//  	),
//  )
//  http.ListenAndServe("localhost:8080", mux)
//
//  -- routing/v1.go --
//
//  func V1(prefix string, handlers ...rest.PackedHandler) (string, http.Handler) {
//  	router := chi.NewRouter()
//  	router.Route(prefix, func(router chi.Router) {
//  		for _, handler := range handlers {
//  			router.Method(handler())
//  		}
//  	})
//  	return prefix, router
//  }
//
//  -- api/v1/pong.go --
//
//  func Pong(path string) rest.Handler {
//  	return func() (string, http.Handler) {
//  		return path, http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) { rw.Write([]byte("pong")) })
//  	}
//  }
//
package rest

import "net/http"

// Handler is a http handler with a specified path.
type Handler func() (path string, handler http.Handler)

// HandlerFunc is a http handler with a specified path.
type HandlerFunc func() (path string, handler http.HandlerFunc)

// PackedHandler is a http handler with a specified http method and path.
type PackedHandler func() (method, path string, handler http.Handler)

// PackedHandlerFunc is a http handler with a specified http method and path.
type PackedHandlerFunc func() (method, path string, handler http.HandlerFunc)
