package rest

import "net/http"

// Routing.
//
//  func V1(prefix string, options ...Option) Handler {
//  	return func() (string, http.Handler) {
//  		cnf := &Configuration{}
//  		for _, configure := range options {
//  			configure(cnf)
//  		}
//
//  		r := chi.NewRouter()
//  		for _, middleware := range cnf.Middlewares {
//  			r.Use(middleware)
//  		}
//
//  		r.Route(prefix, func(r chi.Router) {
//  			for _, handler := range cnf.Handlers {
//  				r.Handle(handler())
//  			}
//  			for _, handler := range cnf.PackedHandlers {
//  				r.Method(handler())
//  			}
//  		})
//
//  		return prefix, r
//  	}
//  }

type Middleware func(http.Handler) http.Handler

type Configuration struct {
	Middlewares    []Middleware
	Handlers       []Handler
	PackedHandlers []PackedHandler
}

type Option func(*Configuration)

func WithMiddlewares(middlewares ...Middleware) Option {
	return func(cnf *Configuration) {
		cnf.Middlewares = append(cnf.Middlewares, middlewares...)
	}
}

func WithHandlers(handlers ...Handler) Option {
	return func(cnf *Configuration) {
		cnf.Handlers = append(cnf.Handlers, handlers...)
	}
}

func WithPackedHandlers(handlers ...PackedHandler) Option {
	return func(cnf *Configuration) {
		cnf.PackedHandlers = append(cnf.PackedHandlers, handlers...)
	}
}
