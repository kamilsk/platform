package context_test

import (
	"context"
	"testing"
	"time"

	. "github.com/kamilsk/platform/internal/sandbox/context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const delta = 10 * time.Millisecond

func TestRun(t *testing.T) {
	delay := func(action func() error, duration time.Duration) func() error {
		return func() error {
			time.Sleep(duration)
			return action()
		}
	}
	canceled := func() context.Context {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		return ctx
	}

	t.Run("serve error", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		srv := server{
			serve:    func() error { return errors.New("serve error") },
			shutdown: func(ctx context.Context) error { return nil },
		}
		assert.EqualError(t, errors.Cause(Run(ctx, srv)), "serve error")
		cancel()
	})
	t.Run("shutdown error", func(t *testing.T) {
		ctx := canceled()
		srv := server{
			serve:    delay(func() error { return errors.New("serve error") }, delta),
			shutdown: func(ctx context.Context) error { return errors.New("shutdown error") },
		}
		assert.EqualError(t, errors.Cause(Run(ctx, srv)), "shutdown error")
	})
	t.Run("context error", func(t *testing.T) {
		chain := Chain(canceled()).Add(canceled())
		srv := server{
			serve:    delay(func() error { return errors.New("serve error") }, delta),
			shutdown: func(ctx context.Context) error { return ctx.Err() },
		}
		assert.Equal(t, errors.Cause(Run(chain, srv)), context.Canceled)
	})
}

type server struct {
	serve    func() error
	shutdown func(context.Context) error
}

func (server server) ListenAndServe() error {
	return server.serve()
}

func (server server) Shutdown(ctx context.Context) error {
	return server.shutdown(ctx)
}
