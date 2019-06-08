package fn_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/pkg/fn"
)

func TestRepeat(t *testing.T) {
	tests := []struct {
		name  string
		times int
	}{
		{"zero times", 0},
		{"constant times", 5},
		{"random times", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000)},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			var counter int
			Repeat(func() { counter++ }, tc.times)
			assert.Equal(t, tc.times, counter)
		})
	}
}
