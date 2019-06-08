package chi

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/kamilsk/platform/protocol/http/rest"
)

// PackHandler packs rest.Handler into rest.PackedHandler using chi router.
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
