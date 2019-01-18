package math_test

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/kamilsk/platform/pkg/math"
	"github.com/stretchr/testify/assert"
)

func TestSequence(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"constant", 5},
		{"random", rand.New(rand.NewSource(time.Now().UnixNano())).Int()},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Len(t, Sequence(tc.size), tc.size)
		})
	}
}
