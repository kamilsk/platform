package sandbox_test

import (
	"context"
	"testing"

	. "github.com/kamilsk/platform/internal/sandbox"
	"github.com/stretchr/testify/assert"
)

func TestFrom(t *testing.T) {
	t.Run("from context", func(t *testing.T) {
		assert.Equal(t, context.TODO(), From(context.Background()).Origin())
	})
	t.Run("from chain", func(t *testing.T) {
		assert.Equal(t, context.TODO(), From(Chain(context.TODO()).Add(context.Background())).Origin())
	})
}
