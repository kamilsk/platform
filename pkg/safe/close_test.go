package safe_test

import (
	"errors"
	"testing"

	. "github.com/kamilsk/platform/pkg/safe"
	"github.com/stretchr/testify/assert"
)

func TestClose(t *testing.T) {
	t.Run("with error", func(t *testing.T) {
		var called bool
		fn := (closer)(func() error { return errors.New("test") })
		Close(fn, func(err error) { called = assert.Error(t, err) })
		assert.True(t, called)
	})
	t.Run("without error", func(t *testing.T) {
		var called bool
		fn := (closer)(func() error { return nil })
		Close(fn, func(error) { called = true })
		assert.False(t, called)
	})
}

type closer func() error

func (fn closer) Close() error {
	return fn()
}
