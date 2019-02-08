package sync_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/kamilsk/platform/pkg/fn"
	. "github.com/kamilsk/platform/pkg/sync"
	"github.com/stretchr/testify/assert"
)

var delta = 10 * time.Millisecond

func TestBreakByDeadline(t *testing.T) {
	t.Run("future deadline", func(t *testing.T) {
		br := BreakByDeadline(time.Now().Add(5 * delta))
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start.Add(5*delta), time.Now(), delta)
	})
	t.Run("passed deadline", func(t *testing.T) {
		br := BreakByDeadline(time.Now().Add(-delta))
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
	t.Run("close multiple times", func(t *testing.T) {
		br := BreakByDeadline(time.Now().Add(time.Hour))
		fn.Repeat(br.Close, 5)
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
}

func TestBreakBySignal(t *testing.T) {
	t.Run("with signal", func(t *testing.T) {
		br := BreakBySignal(os.Interrupt)
		start := time.Now()
		go func() {
			proc, err := os.FindProcess(os.Getpid())
			assert.NoError(t, err)
			assert.NoError(t, proc.Signal(os.Interrupt))
		}()
		<-br.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
	t.Run("without signal", func(t *testing.T) {
		br := BreakBySignal()
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
	t.Run("close multiple times", func(t *testing.T) {
		br := BreakBySignal(os.Kill)
		fn.Repeat(br.Close, 5)
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
}

func TestBreakByTimeout(t *testing.T) {
	t.Run("valid timeout", func(t *testing.T) {
		br := BreakByTimeout(5 * delta)
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start.Add(5*delta), time.Now(), delta)
	})
	t.Run("passed timeout", func(t *testing.T) {
		br := BreakByTimeout(-delta)
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
	t.Run("close multiple times", func(t *testing.T) {
		br := BreakByTimeout(time.Hour)
		fn.Repeat(br.Close, 5)
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
}

func TestMultiplex(t *testing.T) {
	t.Run("with breakers", func(t *testing.T) {
		br := Multiplex(BreakByTimeout(5*delta), BreakByDeadline(time.Now().Add(time.Hour)))
		defer br.Close()
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start.Add(5*delta), time.Now(), delta)
	})
	t.Run("without breakers", func(t *testing.T) {
		br := Multiplex()
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
	t.Run("close multiple times", func(t *testing.T) {
		br := Multiplex(BreakByTimeout(time.Hour))
		fn.Repeat(br.Close, 5)
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
}

func TestMultiplexTwo(t *testing.T) {
	br := MultiplexTwo(
		BreakByDeadline(time.Now().Add(-delta)),
		BreakByTimeout(time.Hour),
	)
	start := time.Now()
	<-br.Done()
	assert.WithinDuration(t, start, time.Now(), delta)
}

func TestMultiplexThree(t *testing.T) {
	br := MultiplexThree(
		BreakByDeadline(time.Now().Add(-delta)),
		BreakBySignal(os.Kill),
		BreakByTimeout(time.Hour),
	)
	start := time.Now()
	<-br.Done()
	assert.WithinDuration(t, start, time.Now(), delta)
}

func TestWithContext(t *testing.T) {
	t.Run("active breaker", func(t *testing.T) {
		ctx := WithContext(context.Background(), BreakByTimeout(5*delta))
		start := time.Now()
		<-ctx.Done()
		assert.WithinDuration(t, start.Add(5*delta), time.Now(), delta)
	})
	t.Run("closed breaker", func(t *testing.T) {
		ctx := WithContext(context.Background(), BreakByTimeout(-delta))
		start := time.Now()
		<-ctx.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
	t.Run("released breaker", func(t *testing.T) {
		br := BreakByTimeout(time.Hour)
		ctx := WithContext(context.Background(), br)
		br.Close()
		start := time.Now()
		<-ctx.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
	t.Run("canceled parent", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		ctx = WithContext(ctx, BreakByTimeout(time.Hour))
		cancel()
		start := time.Now()
		<-ctx.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
}
