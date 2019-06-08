package fn_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/pkg/fn"
)

func TestMust(t *testing.T) {
	tests := []struct {
		name    string
		actions []func() error
		assert  func(assert.TestingT, assert.PanicTestFunc, ...interface{}) bool
	}{
		{
			"with panic",
			[]func() error{
				func() error { return nil },
				func() error { return errors.New("raise panic") },
				func() error { return nil },
			},
			assert.Panics,
		},
		{
			"without panic",
			[]func() error{
				func() error { return nil },
				func() error { return nil },
				func() error { return nil },
			},
			assert.NotPanics,
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			tc.assert(t, func() { Must(tc.actions...) })
		})
	}
}
