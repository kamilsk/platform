package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/protocol/http/middleware"
)

func TestSLA(t *testing.T) {
	var handler http.Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		select {
		case <-time.After(200 * time.Millisecond):
			rw.WriteHeader(http.StatusOK)
		case <-req.Context().Done():
			rw.WriteHeader(http.StatusNoContent)
		}
	})
	t.Run("ok", func(t *testing.T) {
		handler := SLA(300*time.Millisecond, 0.99)(handler)
		rec, req := httptest.NewRecorder(), &http.Request{Method: http.MethodGet}
		handler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, http.StatusOK)
	})
	t.Run("no content", func(t *testing.T) {
		handler := SLA(100*time.Millisecond, 0.99)(handler)
		rec, req := httptest.NewRecorder(), &http.Request{Method: http.MethodGet}
		handler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, http.StatusNoContent)
	})
}
