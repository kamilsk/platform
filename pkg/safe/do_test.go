package safe_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/pkg/safe"
)

func TestSafe(t *testing.T) {
	tests := map[string]struct {
		action  func() error
		handler func(error)
	}{
		"with error": {
			func() error { return errors.New("error") },
			func(err error) { assert.EqualError(t, err, "error") },
		},
		"with panic": {
			func() error { panic("test") },
			func(err error) { assert.EqualError(t, err, `panic unexpected: "test"`) },
		},
		"without anything": {
			func() error { return nil },
			func(error) { t.Fail() },
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) { assert.NotPanics(t, func() { Do(test.action, test.handler) }) })
	}
}
