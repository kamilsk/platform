package chi_test

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/protocol/http/rest"
)

const key = "key"

func v1(prefix string, handlers ...PackedHandler) Handler {
	return func() (string, http.Handler) {
		r := chi.NewRouter()
		r.Route(prefix, func(r chi.Router) {
			for _, handler := range handlers {
				r.Method(handler())
			}
		})
		return prefix, r
	}
}

func v2(prefix string, handlers ...PackedHandlerFunc) HandlerFunc {
	return func() (string, http.HandlerFunc) {
		r := chi.NewRouter()
		r.Route(prefix, func(r chi.Router) {
			for _, handler := range handlers {
				r.Method(handler())
			}
		})
		return prefix, r.ServeHTTP
	}
}

func pingHandler(path string, encoder func(interface{}) ([]byte, error), t assert.TestingT) Handler {
	return func() (string, http.Handler) {
		return pingHandlerFunc(path, encoder, t)()
	}
}

func pingHandlerFunc(path string, encoder func(interface{}) ([]byte, error), t assert.TestingT) HandlerFunc {
	return func() (string, http.HandlerFunc) {
		return path, func(rw http.ResponseWriter, req *http.Request) {
			data, err := encoder(map[string]string{req.URL.Query().Get(key): "pong"})
			assert.NoError(t, err)

			n, err := rw.Write(data)
			assert.NoError(t, err)
			assert.Len(t, data, n)
		}
	}
}
