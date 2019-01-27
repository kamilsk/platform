package sync

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Breaker ...
type Breaker interface {
	Done() <-chan struct{}
	Close()

	trigger() Breaker // guarantee that the done channel will not be nil
}

// BreakByDeadline ...
func BreakByDeadline(deadline time.Time) Breaker {
	timeout := time.Until(deadline)
	if timeout < 0 {
		return closedBreaker()
	}
	return newTimedBreaker(timeout)
}

// BreakBySignal ...
func BreakBySignal(sig ...os.Signal) Breaker {
	if len(sig) == 0 {
		return closedBreaker()
	}
	return newSignaledBreaker(sig)
}

// BreakByTimeout ...
func BreakByTimeout(timeout time.Duration) Breaker {
	if timeout < 0 {
		return closedBreaker()
	}
	return newTimedBreaker(timeout)
}

// WithContext ...
func WithContext(parent context.Context, breaker Breaker) context.Context {
	ctx, cancel := context.WithCancel(parent)
	go func() {
		<-breaker.Done() // this channel is never will be nil, thanks to private trigger() method
		cancel()
	}()
	return ctx
}

func closedBreaker() *breaker {
	br := newBreaker()
	br.Close()
	return br
}

func newBreaker() *breaker {
	return &breaker{signal: make(chan struct{})}
}

type breaker struct {
	signal chan struct{}
	closer sync.Once
}

func (br *breaker) Done() <-chan struct{} {
	return br.signal
}

func (br *breaker) Close() {
	br.closer.Do(func() { close(br.signal) })
}

func (br *breaker) trigger() Breaker {
	return br
}

func newSignaledBreaker(signals []os.Signal) Breaker {
	return (&signaledBreaker{newBreaker(), make(chan os.Signal, len(signals)), signals}).trigger()
}

type signaledBreaker struct {
	*breaker
	relay   chan os.Signal
	signals []os.Signal
}

func (br *signaledBreaker) Close() {
	br.closer.Do(func() {
		signal.Stop(br.relay)
		close(br.signal)
	})
}

func (br *signaledBreaker) trigger() Breaker {
	go func() {
		signal.Notify(br.relay, br.signals...)
		<-br.relay
		br.Close()
	}()
	return br
}

func newTimedBreaker(timeout time.Duration) Breaker {
	return (&timedBreaker{time.NewTimer(timeout), newBreaker()}).trigger()
}

type timedBreaker struct {
	*time.Timer
	*breaker
}

func (br *timedBreaker) Close() {
	br.closer.Do(func() {
		br.Timer.Stop()
		close(br.signal)
	})
}

func (br *timedBreaker) trigger() Breaker {
	go func() {
		<-br.Timer.C
		br.Close()
	}()
	return br
}
