package rest_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/protocol/http/rest"
)

func TestNilMiddlewareFilter(t *testing.T) {
	assert.Nil(t, NilMiddlewareFilter())
	assert.Empty(t, NilMiddlewareFilter(nil, nil, nil))
	assert.Len(t, NilMiddlewareFilter(nil, func(handler http.Handler) http.Handler { return handler }, nil), 1)
}
