package sandbox

import "context"

type ChainedContext interface {
	context.Context

	Add(context.Context) ChainedContext
	Next() ChainedContext
}

func Chain(ctx context.Context) ChainedContext {
	return &chainedContext{Context: ctx}
}

func From(ctx context.Context) ChainedContext {
	if chain, is := ctx.(ChainedContext); is {
		return chain
	}
	return Chain(ctx).Add(context.TODO())
}

type chainedContext struct {
	context.Context
	next *chainedContext
}

func (chain *chainedContext) Add(ctx context.Context) ChainedContext {
	var next, prev = &chain.next, &chain.next
	for *next != nil {
		*prev = *next
		next = &(*next).next
	}
	*prev = &chainedContext{Context: ctx}
	return chain
}

func (chain *chainedContext) Next() ChainedContext {
	return chain.next
}
