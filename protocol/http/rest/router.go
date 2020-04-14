package rest

// RouterConfiguration holds a router configuration.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router/rest instead.
type RouterConfiguration struct {
	Middlewares []Middleware
	Handlers    []Handler
}

// Option applies an option to a router configuration.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router/rest instead.
type Option func(*RouterConfiguration)

// WithMiddlewares adds middlewares to a router configuration.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router/rest instead.
func WithMiddlewares(middlewares ...Middleware) Option {
	return func(cnf *RouterConfiguration) {
		cnf.Middlewares = append(cnf.Middlewares, middlewares...)
	}
}

// WithHandlers adds http handlers with specified paths to a router configuration.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router/rest instead.
func WithHandlers(handlers ...Handler) Option {
	return func(cnf *RouterConfiguration) {
		cnf.Handlers = append(cnf.Handlers, handlers...)
	}
}
