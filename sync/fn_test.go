package sync_test

import (
	"errors"
	"testing"

	. "github.com/kamilsk/platform/sync"
	"github.com/stretchr/testify/assert"
)

func TestSafe(t *testing.T) {
	tests := []struct {
		name   string
		action func() error
		closer func(error)
	}{
		{"with error", func() error { return errors.New("error") },
			func(err error) { assert.EqualError(t, err, "error") }},
		{"with panic", func() error { panic(errors.New("panic")) },
			func(err error) { assert.EqualError(t, err, "unexpected panic handled: panic") }},
		{"without anything", func() error { return nil },
			func(err error) { assert.NoError(t, err) }},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.NotPanics(t, func() { Safe(tc.action, tc.closer) })
		})
	}
}
