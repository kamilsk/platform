package http_test

import (
	"net/http"
	"testing"
	"time"

	. "github.com/kamilsk/platform/internal/sandbox/http"
	"github.com/stretchr/testify/assert"
)

func TestTimeout(t *testing.T) {
	tests := []struct {
		name       string
		header     http.Header
		fallback   time.Duration
		percentage float64
		expected   time.Duration
	}{
		{
			"exists in header",
			http.Header{"X-Timeout": []string{"100ms"}},
			time.Second,
			0.9,
			90 * time.Millisecond,
		},
		{
			"fallback cause empty",
			nil,
			time.Second,
			0.9,
			900 * time.Millisecond,
		},
		{
			"fallback cause invalid",
			http.Header{"X-Timeout": []string{"invalid"}},
			time.Second,
			0.9,
			900 * time.Millisecond,
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Timeout(tc.header, tc.fallback, tc.percentage))
		})
	}
}
