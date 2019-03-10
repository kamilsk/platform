package math_test

import (
	"testing"

	. "github.com/kamilsk/platform/pkg/math"
	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	generator := new(Generator).At(10)

	assert.Equal(t, uint64(10), generator.Current())
	assert.Equal(t, uint64(11), generator.Next())
	assert.Equal(t, uint64(21), generator.Jump(10))
}
