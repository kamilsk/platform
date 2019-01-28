package sync

import (
	"context"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

// A Breaker carries a cancellation signal to break an action execution.
//
// Example based on github.com/kamilsk/retry package:
//
//     if err := retry.Retry(BreakByTimeout(time.Minute), action, strategy.Limit(5)); err != nil {
//             log.Fatal(err)
//     }
//
// Example based on github.com/kamilsk/semaphore package:
//
//     if err := semaphore.Acquire(BreakByTimeout(time.Minute), 5); err != nil {
//             log.Fatal(err)
//     }
//
type Breaker interface {
	// Done returns a channel that's closed when a cancellation signal occurred.
	Done() <-chan struct{}
	// Close closes the Done channel and releases resources associated with it.
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
	closer   sync.Once
	signal   chan struct{}
	released int32
}

// Done returns a channel that's closed when a cancellation signal occurred.
func (br *breaker) Done() <-chan struct{} {
	return br.signal
}

// Close closes the Done channel and releases resources associated with it.
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

// Close closes the Done channel and releases resources associated with it.
func (br *multiplexedBreaker) Close() {
	br.closer.Do(func() {
		each(br.entries).Close()
		close(br.signal)
	})
}

// trigger starts listening all Done channels of multiplexed Breakers.
func (br *multiplexedBreaker) trigger() Breaker {
	go func() {
		brs := make([]reflect.SelectCase, 0, len(br.entries))
		for _, br := range br.entries {
			brs = append(brs, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(br.Done())})
		}
		reflect.Select(brs)
		br.Close()
		atomic.AddInt32(&br.released, 1)
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

// Close closes the Done channel and releases resources associated with it.
func (br *signaledBreaker) Close() {
	br.closer.Do(func() {
		signal.Stop(br.relay)
		close(br.signal)
	})
}

// trigger starts listening required signals to close the Done channel.
func (br *signaledBreaker) trigger() Breaker {
	go func() {
		signal.Notify(br.relay, br.signals...)
		<-br.relay
		br.Close()
		atomic.AddInt32(&br.released, 1)
	}()
	return br
}

func newTimedBreaker(timeout time.Duration) Breaker {
	return &timedBreaker{newBreaker(), time.NewTimer(timeout)}
}

type timedBreaker struct {
	*breaker
	*time.Timer
}

// Close closes the Done channel and releases resources associated with it.
func (br *timedBreaker) Close() {
	br.closer.Do(func() {
		br.Timer.Stop()
		close(br.signal)
	})
}

// trigger starts listening internal timer to close the Done channel.
func (br *timedBreaker) trigger() Breaker {
	go func() {
		<-br.Timer.C
		br.Close()
		atomic.AddInt32(&br.released, 1)
	}()
	return br
}

type each []Breaker

// Close closes all Done channels of a list of Breakers
// and releases resources associated with them.
func (list each) Close() {
	for _, br := range list {
		br.Close()
	}
}
