package chi

import (
	gohttp "net/http"

	"github.com/go-chi/chi"

	"github.com/kamilsk/platform/protocol/http"
	"github.com/kamilsk/platform/protocol/http/rest"
)

// Handler returns handler based on github.com/go-chi/chi.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router/rest instead.
func Handler(prefix string, options ...rest.Option) http.Handler {
	return func() (string, gohttp.Handler) { return Routing(prefix, options...) }
}

// PackHandler packs rest.Handler into rest.Handler using chi router.
//
//  mux := http.NewServeMux()
//  mux.Handle(
//  	chi.Routing("/v1/",
//  		rest.WithPackedHandlers(
//  			chi.PackHandler(gohttp.MethodGet, api.HandlerX),
//  			chi.PackHandler(gohttp.MethodGet, api.HandlerY),
//  			chi.PackHandler(gohttp.MethodGet, api.HandlerZ),
//  		),
//  	),
//  )
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/middleware/chi.Converter instead.
func PackHandler(method string, handler http.Handler, placeholders ...string) rest.Handler {
	if len(placeholders)%2 != 0 {
		panic("count of passed placeholders must be even")
	}
	return func() (string, string, gohttp.Handler) {
		path, h := handler()
		return method, path, gohttp.HandlerFunc(func(rw gohttp.ResponseWriter, req *gohttp.Request) {
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

// PackHandlerFunc packs rest.HandlerFunc into rest.HandlerFunc using chi router.
//
//  mux := http.NewServeMux()
//  mux.HandleFunc(
//  	api.V2("/v2/",
//  		chi.PackHandlerFunc(gohttp.MethodGet, api.HandlerFuncX),
//  		chi.PackHandlerFunc(gohttp.MethodGet, api.HandlerFuncY),
//  		chi.PackHandlerFunc(gohttp.MethodGet, api.HandlerFuncZ),
//  	),
//  )
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/middleware/chi.Converter instead.
func PackHandlerFunc(method string, handler http.HandlerFunc, placeholders ...string) rest.HandlerFunc {
	if len(placeholders)%2 != 0 {
		panic("count of passed placeholders must be even")
	}
	return func() (string, string, gohttp.HandlerFunc) {
		path, h := handler()
		return method, path, func(rw gohttp.ResponseWriter, req *gohttp.Request) {
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

// Routing is a glue for a http listener and the github.com/go-chi/chi router.
//
//  mux := http.NewServeMux()
//  mux.Handle(chi.Routing("/api/", rest.WithMiddlewares(...), rest.WithHandlers(...)))
//  http.ListenAndServe("localhost:8080", mux)
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/router/rest instead.
func Routing(prefix string, options ...rest.Option) (string, gohttp.Handler) {
	cnf := &rest.RouterConfiguration{}
	for _, configure := range options {
		configure(cnf)
	}

	r := chi.NewRouter()
	for _, middleware := range cnf.Middlewares {
		r.Use(middleware)
	}

	r.Route(prefix, func(r chi.Router) {
		for _, handler := range cnf.Handlers {
			r.Method(handler())
		}
	})

	return prefix, r
}
