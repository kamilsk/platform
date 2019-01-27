package sync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreaker_trigger(t *testing.T) {
	br := newBreaker()
	assert.Equal(t, br, br.trigger())
}
