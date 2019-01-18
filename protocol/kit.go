package protocol

import (
	"context"

	"github.com/kamilsk/platform/pkg/sync"
	"github.com/pkg/errors"
)

// Callback contains a channel to return a result of shutdown operation.
type Callback struct {
	context.Context
	Result chan error
}

// Server represents a generic server to listen some protocol.
type Server interface {
	// ListenAndServe listens some protocol and serves it.
	ListenAndServe() error
	// Shutdown tries to do a graceful shutdown.
	Shutdown(context.Context) error
}

// Shutdown is a channel to receive a signal to initiate graceful server shutdown.
type Shutdown chan Callback

// Run runs ListenAndServe in separated goroutine and listens shutdown signal.
// It returns ListenAndServe' error or Shutdown' error if signal is received.
func Run(server Server, shutdown Shutdown) error {
	serve := make(chan error, 1)

	go sync.Safe(func() error {
		return server.ListenAndServe()
	}, func(err error) {
		serve <- errors.Wrap(err, "tried to listen and serve a connection")
		close(serve)
	})

	select {
	case callback := <-shutdown:
		shutdownErr := errors.Wrap(server.Shutdown(callback.Context), "tried to shutting down the server")
		callback.Result <- shutdownErr
		return shutdownErr
	case serveErr := <-serve:
		return serveErr
	}
}
