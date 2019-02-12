package signal_test

import (
	"context"
	"os"
	"testing"

	. "github.com/kamilsk/platform/protocol/signal"
	"github.com/stretchr/testify/assert"
)

func TestListener_Listen(t *testing.T) {
	t.Run("break by context", func(t *testing.T) {
		listener := New()
		listener.Callback(func() { t.Fail() })
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		assert.NoError(t, listener.Listen(ctx))
	})
	t.Run("break by signal", func(t *testing.T) {
		var success bool
		listener := New()
		listener.Callback(func() { success = true })
		go func() {
			proc, err := os.FindProcess(os.Getpid())
			assert.NoError(t, err)
			assert.NoError(t, proc.Signal(os.Interrupt))
		}()
		assert.NoError(t, listener.Listen(context.Background()))
		assert.True(t, success)
	})
}
