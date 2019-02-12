package signal

import (
	"context"

	"github.com/kamilsk/platform/pkg/sync"
)

// New returns a listener of the termination signals.
func New() *listener {
	return &listener{}
}

type listener struct {
	callbacks []func()
}

// Callback registers a callback to execution later
// when the termination signals caught.
func (listener *listener) Callback(callback func()) {
	listener.callbacks = append(listener.callbacks, callback)
}

// Listen starts listening to termination signals.
// It runs registered callbacks when the termination signals caught
// and never returns an error.
func (listener *listener) Listen(ctx context.Context) error {
	if err := sync.Termination().Wait(ctx); err == sync.ErrSignalTrapped {
		for _, callback := range listener.callbacks {
			callback()
		}
	}
	return nil
}
