package sync_test

import (
	"context"
	"os"
	"testing"
	"time"

	. "github.com/kamilsk/platform/pkg/sync"
	"github.com/stretchr/testify/assert"
)

func TestTermination(t *testing.T) {
	tests := []struct {
		name     string
		breaker  func(cancel context.CancelFunc)
		expected error
	}{
		{
			"break by signal",
			func(cancel context.CancelFunc) {
				proc, err := os.FindProcess(os.Getpid())
				assert.NoError(t, err)
				assert.NoError(t, proc.Signal(os.Interrupt))
				time.Sleep(delta)
				cancel()
			},
			ErrSignalTrapped,
		},
		{
			"break by context",
			func(cancel context.CancelFunc) {
				cancel()
			},
			context.Canceled,
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			trap := Termination()
			go tc.breaker(cancel)
			assert.Equal(t, tc.expected, trap.Wait(ctx))
		})
	}
}
