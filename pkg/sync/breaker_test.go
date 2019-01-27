package sync_test

import (
	"context"
	"os"
	"testing"
	"time"

	. "github.com/kamilsk/platform/pkg/sync"
	"github.com/stretchr/testify/assert"
)

var delta = 10 * time.Millisecond

func TestBreakByDeadline(t *testing.T) {
	t.Run("future", func(t *testing.T) {
		br := BreakByDeadline(time.Now().Add(5 * delta))
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start.Add(5*delta), time.Now(), delta)
	})
	t.Run("past", func(t *testing.T) {
		br := BreakByDeadline(time.Now().Add(-delta))
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
}

func TestBreakByTimeout(t *testing.T) {
	t.Run("future", func(t *testing.T) {
		br := BreakByTimeout(5 * delta)
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start.Add(5*delta), time.Now(), delta)
	})
	t.Run("past", func(t *testing.T) {
		br := BreakByTimeout(-delta)
		start := time.Now()
		<-br.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
}

func TestWithContext(t *testing.T) {
	t.Run("future", func(t *testing.T) {
		br := BreakByTimeout(5 * delta)
		ctx := WithContext(context.Background(), br)
		start := time.Now()
		<-ctx.Done()
		assert.WithinDuration(t, start.Add(5*delta), time.Now(), delta)
	})
	t.Run("past", func(t *testing.T) {
		br := BreakByTimeout(-delta)
		ctx := WithContext(context.Background(), br)
		start := time.Now()
		<-ctx.Done()
		assert.WithinDuration(t, start, time.Now(), delta)
	})
}
