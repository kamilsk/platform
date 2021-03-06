package chi_test

import (
	"encoding/json"
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kamilsk/platform/protocol/http/rest"
	. "github.com/kamilsk/platform/protocol/http/rest/router/chi"
)

func TestPackHandler(t *testing.T) {
	t.Run("normal case", func(t *testing.T) {
		api := "/v1/"
		mux := http.NewServeMux()
		apiHandler := Handler(
			api,
			rest.WithMiddlewares(
				rest.NilMiddlewareFilter(
					nil,
					func(handler http.Handler) http.Handler { return handler },
					nil,
				)...,
			),
			rest.WithHandlers(
				PackHandler(http.MethodGet, pingHandler("/{id}", json.Marshal, t), "id", key),
				PackHandler(http.MethodGet, welcomeHandler("/{greeting}/{person}", t), "greeting", welcome, "person", name),
			),
		)
		mux.Handle(apiHandler())
		assert.JSONEq(t, `{"test":"pong"}`, assert.HTTPBody(mux.ServeHTTP, http.MethodGet, path.Join(api, "test"), nil))
		assert.Equal(t, "Hello, World!", assert.HTTPBody(mux.ServeHTTP, http.MethodGet, path.Join(api, "Hello", "World"), nil))
	})
	t.Run("invalid placeholders", func(t *testing.T) {
		assert.Panics(t, func() { PackHandler(http.MethodGet, welcomeHandler("/{greeting}/{person}", t), "greeting") })
	})
}

func TestPackHandlerFunc(t *testing.T) {
	t.Run("normal case", func(t *testing.T) {
		api := "/v2/"
		mux := http.NewServeMux()
		mux.HandleFunc(v2(api,
			PackHandlerFunc(http.MethodGet, pingHandlerFunc("/{id}", json.Marshal, t), "id", key),
			PackHandlerFunc(http.MethodGet, welcomeHandlerFunc("/{greeting}/{person}", t), "greeting", welcome, "person", name),
		))
		assert.JSONEq(t, `{"test":"pong"}`, assert.HTTPBody(mux.ServeHTTP, http.MethodGet, path.Join(api, "test"), nil))
		assert.Equal(t, "Hello, World!", assert.HTTPBody(mux.ServeHTTP, http.MethodGet, path.Join(api, "Hello", "World"), nil))
	})
	t.Run("invalid placeholders", func(t *testing.T) {
		assert.Panics(t, func() { PackHandlerFunc(http.MethodGet, welcomeHandlerFunc("/{greeting}/{person}", t), "greeting") })
	})
}
