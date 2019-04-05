package context

import "context"

type ClosableContext interface {
	context.Context
	Close()
}

func Closable(ctx context.Context, cancel context.CancelFunc) ClosableContext {
	return &closable{ctx, cancel}
}

type closable struct {
	context.Context
	cancel context.CancelFunc
}

func (closable *closable) Close() {
	closable.cancel()
}
