package io_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/kamilsk/platform/pkg/io"
	"github.com/stretchr/testify/assert"
)

func TestTeeReadCloser(t *testing.T) {
	payload := "invalid json"

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		buf := bytes.NewBuffer(nil)
		body := TeeReadCloser(req.Body, buf)

		var expected []int
		assert.Error(t, json.NewDecoder(body).Decode(&expected))
		assert.Equal(t, payload, buf.String())
	}))
	defer server.Close()

	resp, err := http.Post(server.URL, "application/json", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
