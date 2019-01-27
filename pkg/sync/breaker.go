package sync

import (
	"context"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"time"
)

// A Breaker carries a cancellation signal to break an action execution.
//
// Example based on github.com/kamilsk/retry package:
//
//     if err := retry.Retry(sync.BreakByTimeout(time.Minute), action, strategy.Limit(5)); err != nil {
//             log.Fatal(err)
//     }
//
// Example based on github.com/kamilsk/semaphore package:
//
//     if err := semaphore.Acquire(sync.BreakByTimeout(time.Minute), 5); err != nil {
//             log.Fatal(err)
//     }
//
type Breaker interface {
	// Done returns a channel that's closed when a cancellation signal occurred.
	Done() <-chan struct{}
	// Close closes the done channel and releases resources associated with it.
	Close()
	// trigger is a private method to guarantee that the breakers come from
	// this package and all of them return a valid done channel.
	trigger() Breaker
}

// BreakByDeadline ...
func BreakByDeadline(deadline time.Time) Breaker {
	timeout := time.Until(deadline)
	if timeout < 0 {
		return closedBreaker()
	}
	return newTimedBreaker(timeout).trigger()
}

// BreakBySignal ...
func BreakBySignal(sig ...os.Signal) Breaker {
	if len(sig) == 0 {
		return closedBreaker()
	}
	return newSignaledBreaker(sig).trigger()
}

// BreakByTimeout ...
func BreakByTimeout(timeout time.Duration) Breaker {
	if timeout < 0 {
		return closedBreaker()
	}
	return newTimedBreaker(timeout).trigger()
}

// Multiplex ...
func Multiplex(breakers ...Breaker) Breaker {
	if len(breakers) == 0 {
		return closedBreaker()
	}
	return newMultiplexedBreaker(breakers).trigger()
}

// WithContext ...
func WithContext(parent context.Context, breaker Breaker) context.Context {
	ctx, cancel := context.WithCancel(parent)
	go func() {
		<-breaker.Done()
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

func newMultiplexedBreaker(entries []Breaker) Breaker {
	return &multiplexedBreaker{newBreaker(), entries}
}

type multiplexedBreaker struct {
	*breaker
	entries []Breaker
}

func (br *multiplexedBreaker) Close() {
	br.closer.Do(func() {
		each(br.entries).Close()
		close(br.signal)
	})
}

func (br *multiplexedBreaker) trigger() Breaker {
	go func() {
		brs := make([]reflect.SelectCase, 0, len(br.entries))
		for _, br := range br.entries {
			brs = append(brs, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(br.Done())})
		}
		reflect.Select(brs)
		br.Close()
	}()
	return br
}

func newSignaledBreaker(signals []os.Signal) Breaker {
	return &signaledBreaker{newBreaker(), make(chan os.Signal, len(signals)), signals}
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
	return &timedBreaker{time.NewTimer(timeout), newBreaker()}
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

type each []Breaker

func (list each) Close() {
	for _, br := range list {
		br.Close()
	}
}
