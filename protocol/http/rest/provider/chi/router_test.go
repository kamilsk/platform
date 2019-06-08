package chi_test

import (
	"encoding/json"
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/protocol/http/rest/provider/chi"
)

func TestPackHandler(t *testing.T) {
	api := "/v1/"
	mux := http.NewServeMux()
	mux.Handle(v1(api,
		PackHandler(http.MethodGet, pingHandler("/{id}", json.Marshal, t), "id", key),
		PackHandler(http.MethodGet, welcomeHandler("/{greeting}/{person}", t), "greeting", welcome, "person", name),
	)())
	assert.JSONEq(t, `{"test":"pong"}`, assert.HTTPBody(mux.ServeHTTP, http.MethodGet, path.Join(api, "test"), nil))
	assert.Equal(t, "Hello, World!", assert.HTTPBody(mux.ServeHTTP, http.MethodGet, path.Join(api, "Hello", "World"), nil))
}

func TestPackHandlerFunc(t *testing.T) {
	api := "/v2/"
	mux := http.NewServeMux()
	mux.Handle(v2(api,
		PackHandlerFunc(http.MethodGet, pingHandlerFunc("/{id}", json.Marshal, t), "id", key),
		PackHandlerFunc(http.MethodGet, welcomeHandlerFunc("/{greeting}/{person}", t), "greeting", welcome, "person", name),
	)())
	assert.JSONEq(t, `{"test":"pong"}`, assert.HTTPBody(mux.ServeHTTP, http.MethodGet, path.Join(api, "test"), nil))
	assert.Equal(t, "Hello, World!", assert.HTTPBody(mux.ServeHTTP, http.MethodGet, path.Join(api, "Hello", "World"), nil))
}
