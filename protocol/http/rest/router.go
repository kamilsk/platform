package rest

import "net/http"

// RouterConfiguration holds a router configuration.
type RouterConfiguration struct {
	Middlewares []Middleware
	Handlers    []Handler
}

// Middleware provides functionality to compose http handlers.
type Middleware func(http.Handler) http.Handler

// Option applies an option to a router configuration.
type Option func(*RouterConfiguration)

// WithMiddlewares adds middlewares to a router configuration.
func WithMiddlewares(middlewares ...Middleware) Option {
	return func(cnf *RouterConfiguration) { cnf.Middlewares = append(cnf.Middlewares, middlewares...) }
}

// WithHandlers adds http handlers with specified paths to a router configuration.
func WithHandlers(handlers ...Handler) Option {
	return func(cnf *RouterConfiguration) { cnf.Handlers = append(cnf.Handlers, handlers...) }
}
