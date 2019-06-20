package rest

import "net/http"

// Configuration holds a router configuration.
type Configuration struct {
	Middlewares    []Middleware
	Handlers       []Handler
	PackedHandlers []PackedHandler
}

// Middleware provides functionality to compose http handlers.
type Middleware func(http.Handler) http.Handler

// Option applies an option to a router configuration.
type Option func(*Configuration)

// WithMiddlewares adds middlewares to a router configuration.
func WithMiddlewares(middlewares ...Middleware) Option {
	return func(cnf *Configuration) { cnf.Middlewares = append(cnf.Middlewares, middlewares...) }
}

// WithHandlers adds http handlers with specified paths to a router configuration.
func WithHandlers(handlers ...Handler) Option {
	return func(cnf *Configuration) { cnf.Handlers = append(cnf.Handlers, handlers...) }
}

// WithPackedHandlers adds http handlers with specified http methods and paths to a router configuration.
func WithPackedHandlers(handlers ...PackedHandler) Option {
	return func(cnf *Configuration) { cnf.PackedHandlers = append(cnf.PackedHandlers, handlers...) }
}
