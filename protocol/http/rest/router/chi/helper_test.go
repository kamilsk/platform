package chi_test

import (
	"fmt"
	gohttp "net/http"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"

	"github.com/kamilsk/platform/protocol/http"
	. "github.com/kamilsk/platform/protocol/http/rest"
)

const (
	key     = "key"
	name    = "name"
	welcome = "welcome"
)

func v2(prefix string, handlers ...HandlerFunc) (string, gohttp.HandlerFunc) {
	router := chi.NewRouter()
	router.Route(prefix, func(r chi.Router) {
		for _, handler := range handlers {
			r.Method(handler())
		}
	})
	return prefix, router.ServeHTTP
}

func pingHandler(path string, encoder func(interface{}) ([]byte, error), t assert.TestingT) http.Handler {
	return func() (string, gohttp.Handler) { return pingHandlerFunc(path, encoder, t)() }
}

func pingHandlerFunc(path string, encoder func(interface{}) ([]byte, error), t assert.TestingT) http.HandlerFunc {
	return func() (string, gohttp.HandlerFunc) {
		return path, func(rw gohttp.ResponseWriter, req *gohttp.Request) {
			data, err := encoder(map[string]string{req.URL.Query().Get(key): "pong"})
			assert.NoError(t, err)

			n, err := rw.Write(data)
			assert.NoError(t, err)
			assert.Len(t, data, n)
		}
	}
}

func welcomeHandler(path string, t assert.TestingT) http.Handler {
	return func() (string, gohttp.Handler) { return welcomeHandlerFunc(path, t)() }
}

func welcomeHandlerFunc(path string, t assert.TestingT) http.HandlerFunc {
	return func() (string, gohttp.HandlerFunc) {
		return path, func(rw gohttp.ResponseWriter, req *gohttp.Request) {
			q := req.URL.Query()

			n, err := fmt.Fprintf(rw, "%s, %s!", q.Get(welcome), q.Get(name))
			assert.NoError(t, err)
			assert.True(t, n > 5)
		}
	}
}
