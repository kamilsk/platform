package safe_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/pkg/safe"
)

func TestSafe(t *testing.T) {
	tests := []struct {
		name   string
		action func() error
		closer func(error)
	}{
		{
			"with error",
			func() error { return errors.New("error") },
			func(err error) { assert.EqualError(t, err, "error") },
		},
		{
			"with panic",
			func() error { panic("test") },
			func(err error) { assert.EqualError(t, err, "unexpected panic handled: test") },
		},
		{
			"without anything",
			func() error { return nil },
			func(err error) { assert.NoError(t, err) },
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.NotPanics(t, func() { Do(tc.action, tc.closer) })
		})
	}
}
