package sandbox

import (
	"context"

	"github.com/kamilsk/platform/pkg/safe"
	"github.com/kamilsk/platform/protocol"
	"github.com/pkg/errors"
)

func Run(ctx context.Context, server protocol.Server) error {
	serve := make(chan error, 1)

	go safe.Do(func() error {
		return server.ListenAndServe()
	}, func(err error) {
		serve <- errors.Wrap(err, "tried to listen and serve a connection")
		close(serve)
	})

	select {
	case <-ctx.Done():
		return errors.Wrap(server.Shutdown(From(ctx).Next()), "tried to shutting down the server")
	case err := <-serve:
		return err
	}
}
