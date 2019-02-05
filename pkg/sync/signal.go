package sync

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
)

// ErrSignalTrapped is returned by the SignalTrap.Wait
// when the expected signals caught.
var ErrSignalTrapped = errors.New("signal trapped")

// Termination returns trap for termination signals.
func Termination() SignalTrap {
	trap := make(chan os.Signal, 3)
	signal.Notify(trap, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	return trap
}

// SignalTrap wraps os.Signal channel to provide some useful methods.
//
//  trap := make(chan SignalTrap)
//  signal.Notify(trap, os.Interrupt)
//  trap.Wait(context.Background())
//
type SignalTrap chan os.Signal

// Wait blocks until one of the expected signals caught
// or the Context closed. It unregisters from the notification
// and closes itself.
func (trap SignalTrap) Wait(ctx context.Context) error {
	defer close(trap)
	defer signal.Stop(trap)
	select {
	case <-trap:
		return ErrSignalTrapped
	case <-ctx.Done():
		return ctx.Err()
	}
}
