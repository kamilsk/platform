package fn_test

import (
	"testing"
	"time"

	. "github.com/kamilsk/platform/pkg/fn"
	"github.com/stretchr/testify/assert"
)

var delta = 10 * time.Millisecond

func TestStopwatch(t *testing.T) {
	var compare time.Duration

	duration := Stopwatch(func() {
		start := time.Now()
		time.Sleep(delta)
		compare = time.Since(start)
	})

	assert.True(t, compare < duration)
}
