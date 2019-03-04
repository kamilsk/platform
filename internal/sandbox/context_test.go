package sandbox_test

import (
	"context"
	"testing"

	. "github.com/kamilsk/platform/internal/sandbox"
	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	validate := func(chain ChainedContext) {
		assert.Equal(t, context.TODO(), chain.Origin())
		assert.NotNil(t, chain.Next())
		assert.NotEqual(t, context.TODO(), chain.Next().Origin())
		assert.NotNil(t, chain.Next().Next())
		assert.NotEqual(t, context.TODO(), chain.Next().Next().Origin())
		assert.Nil(t, chain.Next().Next().Next())
	}

	t.Run("simple", func(t *testing.T) {
		chain := Chain(context.TODO()).
			Add(context.WithValue(context.TODO(), "key1", "value1")).
			Add(context.WithValue(context.TODO(), "key2", "value2"))
		validate(chain)
	})
	t.Run("complex", func(t *testing.T) {
		chain := Chain(context.TODO())
		chain.
			Add(context.WithValue(context.TODO(), "key1", "value1")).
			Next().
			Add(context.WithValue(context.TODO(), "key2", "value2"))
		validate(chain)
	})
}

func TestFrom(t *testing.T) {
	t.Run("from context", func(t *testing.T) {
		assert.Equal(t, context.TODO(), From(context.TODO()).Origin())
	})
	t.Run("from chain", func(t *testing.T) {
		assert.Equal(t, context.TODO(), From(Chain(context.TODO()).Add(context.Background())).Origin())
	})
}
