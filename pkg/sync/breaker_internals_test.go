package sync

import (
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBreaker_trigger(t *testing.T) {
	br := newBreaker()
	assert.Equal(t, br, br.trigger())
}

func TestMultiplexedBreaker_Close(t *testing.T) {
	br := Multiplex(BreakBySignal(os.Kill), BreakByTimeout(time.Hour)).(*multiplexedBreaker)
	br.Close()
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, int32(1), atomic.LoadInt32(&br.released))
}

func TestSignaledBreaker_Close(t *testing.T) {
	br := BreakBySignal(os.Kill).(*signaledBreaker)
	br.Close()
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, int32(1), atomic.LoadInt32(&br.released))
}

func TestTimedBreaker_Close(t *testing.T) {
	br := BreakByTimeout(time.Hour).(*timedBreaker)
	br.Close()
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, int32(1), atomic.LoadInt32(&br.released))
}
