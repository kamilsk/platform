package context

import "context"

type ChainedContext interface {
	context.Context

	Add(func() context.Context) ChainedContext
	Next() ChainedContext
	Origin() context.Context
}

func Chain(ctx context.Context) ChainedContext {
	return &node{Context: ctx}
}

func From(ctx context.Context) ChainedContext {
	if chain, is := ctx.(ChainedContext); is {
		return chain
	}
	return Chain(ctx).Add(func() context.Context { return Closable(context.WithCancel(context.Background())) })
}

type node struct {
	context.Context
	builder func() context.Context
	next    *node
}

func (chain *node) Add(next func() context.Context) ChainedContext {
	link := &chain.next
	for *link != nil {
		link = &(*link).next
	}
	*link = &node{builder: next}
	return chain
}

func (chain *node) Next() ChainedContext {
	return chain.next
}

func (chain *node) Origin() context.Context {
	if chain.Context == nil {
		chain.Context = chain.builder()
	}
	return chain.Context
}
