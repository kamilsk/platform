package sandbox

import "context"

type ChainedContext interface {
	context.Context

	Add(context.Context) ChainedContext
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
	return Chain(ctx).Add(context.TODO())
}

type node struct {
	context.Context
	next *node
}

func (chain *node) Add(ctx context.Context) ChainedContext {
	next := &chain.next
	for *next != nil {
		next = &(*next).next
	}
	*next = &node{Context: ctx}
	return chain
}

func (chain *node) Next() ChainedContext {
	return chain.next
}

func (chain *node) Origin() context.Context {
	return chain.Context
}
