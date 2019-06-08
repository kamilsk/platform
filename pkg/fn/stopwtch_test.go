package fn_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/pkg/fn"
)

const delta = 10 * time.Millisecond

func TestStopwatch(t *testing.T) {
	var compare time.Duration

	duration := Stopwatch(func() {
		start := time.Now()
		time.Sleep(delta)
		compare = time.Since(start)
	})

	assert.True(t, compare < duration)
}
