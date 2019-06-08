package signal_test

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kamilsk/platform/pkg/safe"
	. "github.com/kamilsk/platform/protocol/signal"
)

func TestListener_Listen(t *testing.T) {
	t.Run("break by context", func(t *testing.T) {
		listener := New()
		listener.AddResource(safe.Releaser(func() error { panic("unexpected") }))
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		assert.NoError(t, listener.Listen(ctx))
	})
	t.Run("break by signal", func(t *testing.T) {
		var success bool
		listener := New()
		listener.AddResource(
			safe.Releaser(func() error { return errors.New("test") }),
			func(err error) { success = err != nil && strings.Contains(err.Error(), "test") },
		)
		go func() {
			proc, err := os.FindProcess(os.Getpid())
			assert.NoError(t, err)
			assert.NoError(t, proc.Signal(os.Interrupt))
		}()
		assert.NoError(t, listener.Listen(context.Background()))
		assert.True(t, success)
	})
}
