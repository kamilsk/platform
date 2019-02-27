package sandbox

import (
	"context"

	"github.com/kamilsk/platform/pkg/safe"
	"github.com/kamilsk/platform/protocol"
	"github.com/pkg/errors"
)

type ChainedContext interface {
	context.Context

	Add(context.Context) ChainedContext
	Next() ChainedContext
}

func New(ctx context.Context) ChainedContext {
	return &chainedContext{Context: ctx}
}

type chainedContext struct {
	context.Context
	next *chainedContext
}

func (chain *chainedContext) Add(ctx context.Context) ChainedContext {
	var next, prev = &chain.next, &chain.next
	for *next != nil {
		*prev = *next
	}
	*prev = &chainedContext{Context: ctx}
	return chain
}

func (chain *chainedContext) Next() ChainedContext {
	return chain.next
}

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
		if chain, is := ctx.(ChainedContext); is {
			ctx = chain.Next()
		}
		shutdownErr := errors.Wrap(server.Shutdown(ctx), "tried to shutting down the server")
		return shutdownErr
	case serveErr := <-serve:
		return serveErr
	}
}
