package protocol

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kamilsk/platform/pkg/safe"
)

// Callback contains a channel to return a result of shutdown operation.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/server instead.
type Callback struct {
	context.Context
	Result chan error
}

// Shutdown is a channel to receive a signal to initiate graceful server shutdown.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/server instead.
type Shutdown chan Callback

// Run runs ListenAndServe in separated goroutine and listens shutdown signal.
// It returns ListenAndServe' error or Shutdown' error if signal is received.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/server.Run instead.
func Run(server Server, shutdown Shutdown) error {
	serve := make(chan error, 1)

	go safe.Do(func() error {
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
