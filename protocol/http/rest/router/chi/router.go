package chi

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/kamilsk/platform/protocol/http/rest"
)

// PackHandler packs rest.Handler into rest.PackedHandler using chi router.
//
//  mux := http.NewServeMux()
//  mux.Handle(
//  	rest.V1("/v1/",
//  		chi.PackHandler(gohttp.MethodGet, api.HandlerX),
//  		chi.PackHandler(gohttp.MethodGet, api.HandlerY),
//  		chi.PackHandler(gohttp.MethodGet, api.HandlerZ),
//  	),
//  )
//
func PackHandler(method string, handler rest.Handler, placeholders ...string) rest.PackedHandler {
	if len(placeholders)%2 != 0 {
		panic("count of passed placeholders must be even")
	}
	return func() (string, string, http.Handler) {
		path, h := handler()
		return method, path, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if steps := len(placeholders); steps > 0 {
				q := req.URL.Query()
				for i := 0; i < steps; i += 2 {
					from, to := placeholders[i], placeholders[i+1]
					q.Set(to, chi.URLParam(req, from))
				}
				req.URL.RawQuery = q.Encode()
			}
			h.ServeHTTP(rw, req)
		})
	}
}

// PackHandlerFunc packs rest.HandlerFunc into rest.PackedHandlerFunc using chi router.
//
//  mux := http.NewServeMux()
//  mux.HandleFunc(
//  	rest.V2("/v2/",
//  		chi.PackHandlerFunc(gohttp.MethodGet, api.HandlerFuncX),
//  		chi.PackHandlerFunc(gohttp.MethodGet, api.HandlerFuncY),
//  		chi.PackHandlerFunc(gohttp.MethodGet, api.HandlerFuncZ),
//  	),
//  )
//
func PackHandlerFunc(method string, handler rest.HandlerFunc, placeholders ...string) rest.PackedHandlerFunc {
	if len(placeholders)%2 != 0 {
		panic("count of passed placeholders must be even")
	}
	return func() (string, string, http.HandlerFunc) {
		path, h := handler()
		return method, path, func(rw http.ResponseWriter, req *http.Request) {
			if steps := len(placeholders); steps > 0 {
				q := req.URL.Query()
				for i := 0; i < steps; i += 2 {
					from, to := placeholders[i], placeholders[i+1]
					q.Set(to, chi.URLParam(req, from))
				}
				req.URL.RawQuery = q.Encode()
			}
			h.ServeHTTP(rw, req)
		}
	}
}

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
