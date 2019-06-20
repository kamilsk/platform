package rest_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/protocol/http/rest"
)

func TestRouterConfiguration(t *testing.T) {
	t.Run("with middlewares", func(t *testing.T) {
		cnf := &RouterConfiguration{}
		configure := WithMiddlewares(func(handler http.Handler) http.Handler { return handler })
		configure(cnf)

		assert.Len(t, cnf.Middlewares, 1)
		assert.Empty(t, cnf.Handlers)
		assert.Empty(t, cnf.PackedHandlers)
	})
	t.Run("with handlers", func(t *testing.T) {
		cnf := &RouterConfiguration{}
		configure := WithHandlers(func() (string, http.Handler) { return "/", http.DefaultServeMux })
		configure(cnf)

		assert.Len(t, cnf.Handlers, 1)
		assert.Empty(t, cnf.Middlewares)
		assert.Empty(t, cnf.PackedHandlers)
	})
	t.Run("with packed handlers", func(t *testing.T) {
		cnf := &RouterConfiguration{}
		configure := WithPackedHandlers(func() (string, string, http.Handler) {
			return http.MethodGet, "/", http.DefaultServeMux
		})
		configure(cnf)

		assert.Len(t, cnf.PackedHandlers, 1)
		assert.Empty(t, cnf.Middlewares)
		assert.Empty(t, cnf.Handlers)
	})
}
